package services

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	internalmodels "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/structs"
	"go.uber.org/zap"
)

type Project struct {
	logger      *zap.Logger
	projectRepo repos.Project
	teamRepo    repos.Team
}

// NewProject returns a new Project service.
func NewProject(logger *zap.Logger,
	projectRepo repos.Project,
	teamRepo repos.Team,
) *Project {
	return &Project{
		logger:      logger,
		projectRepo: projectRepo,
		teamRepo:    teamRepo,
	}
}

// addTeam, removeTeam, projectByID
func (p *Project) ProjectByID(ctx context.Context, d db.DBTX, projectID int) (*db.Project, error) {
	defer newOTELSpan(ctx, "Project.ProjectByID").End()

	project, err := p.projectRepo.ProjectByID(ctx, d, projectID)
	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeNotFound, "project not found")
	}

	return project, nil
}

// TODO accept array of paths to initialize obj with (this will be done at app startup right before running server and update db).
// therefore config is always up to date in the backend. obj2 will be empty
// when a user updates the config in UI we use this same function but with empty [] of paths
// therefore we use db's config untouched as is.
// merging logic is the same for both scenarios.
// obj1 : required existing database config
// obj2 : optional user config update request
// pathKeys : optional keys to generate the initial config from. obj1 will be merged into this object
// TODO accepts projectID to get both pathKeys and obj1 every time
// (we dont know any of those, just projectID)
func (p *Project) MergeConfigFields(ctx context.Context, d db.DBTX, projectID int, obj2 map[string]any) (*models.ProjectConfig, error) {
	project, err := p.projectRepo.ProjectByID(ctx, d, projectID)
	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeNotFound, "project not found")
	}

	fieldsMap := make(map[string]map[string]any)

	var obj1 models.ProjectConfig
	err = json.Unmarshal(project.BoardConfig.Bytes, &obj1)
	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeUnknown, "invalid stored board config: %v", err)
	}

	// fmt.Printf("project.BoardConfig: %v\n", string(project.BoardConfig.Bytes))
	fmt.Printf("obj1: %v\n", obj1)

	var workItem any
	switch internalmodels.Project(project.Name) {
	case internalmodels.ProjectDemoProject:
		// FIXME due to structs package, its broken if all fields are pointers
		// because they're all nil reflect.Pointer's so no keys are generated
		// either put shared models in a literal "internalmodels" folder
		// and let generated model types be or fix oapicodegen or
		// initialize the below struct with zero values
		workItem := &internalmodels.RestDemoProjectWorkItemsResponse{}
		createZeroStructFields(reflect.ValueOf(workItem), 10)
		format.PrintJSON(workItem)
	}
	pathKeys := structs.GetKeys(workItem, "")
	fmt.Printf("pathKeys: %v\n", pathKeys)

	for _, path := range pathKeys {
		fieldsMap[path] = defaultConfigField(path)
	}

	var obj1Map map[string]any
	fj, _ := json.Marshal(obj1)
	json.Unmarshal(fj, &obj1Map)

	p.mergeFieldsMap(fieldsMap, obj1Map)

	p.mergeFieldsMap(fieldsMap, obj2)

	obj1.Fields = make([]models.ProjectConfigField, 0, len(fieldsMap))
	for _, field := range fieldsMap {
		var fieldStruct models.ProjectConfigField

		fBlob, _ := json.Marshal(field)
		_ = json.Unmarshal(fBlob, &fieldStruct)

		obj1.Fields = append(obj1.Fields, fieldStruct)
	}

	return &obj1, err
}

func defaultConfigField(path string) map[string]any {
	f := models.ProjectConfigField{
		Path:          path,
		Name:          path[strings.LastIndex(path, ".")+1:],
		IsVisible:     true,
		IsEditable:    true,
		ShowCollapsed: true,
	}

	var jsonMap map[string]any

	fj, _ := json.Marshal(f)
	_ = json.Unmarshal(fj, &jsonMap)

	return jsonMap
}

// https://github.com/icza/dyno looks promising
func (p *Project) mergeFieldsMap(fieldsMap map[string]map[string]any, obj map[string]any) {
	fieldsInterface, ok := obj["fields"]
	if !ok {
		return
	}
	var fields []map[string]any // can't type assert map values of any when obj comes from unmarshalling
	fBlob, err := json.Marshal(fieldsInterface)
	if err != nil {
		return
	}
	if err := json.Unmarshal(fBlob, &fields); err != nil {
		return
	}

	for _, fieldMap := range fields {
		path, ok := fieldMap["path"].(string)
		if !ok {
			continue
		}

		if fm, ok := fieldsMap[path]; ok {
			newField := defaultConfigField(fm["path"].(string))

			for key, value := range fieldMap {
				if reflect.TypeOf(value) != reflect.TypeOf(newField[key]) {
					continue
				}
				fmt.Printf("value: %v\n", value)
				newField[key] = value
			}
			fieldsMap[path] = newField
		}
	}
}

func createZeroStructFields(v reflect.Value, maxDepth int) {
	if maxDepth == 0 {
		return
	}
	maxDepth--
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		createZeroStructFields(v.Elem(), maxDepth)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			if field.CanSet() {
				createZeroStructFields(field, maxDepth)
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			createZeroStructFields(v.Index(i), maxDepth)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			createZeroStructFields(v.MapIndex(key), maxDepth)
		}
	}
}
