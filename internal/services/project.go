package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/structs"
	"go.uber.org/zap"
)

type Project struct {
	logger      *zap.SugaredLogger
	projectRepo repos.Project
	teamRepo    repos.Team
}

// NewProject returns a new Project service.
func NewProject(logger *zap.SugaredLogger, projectRepo repos.Project,
	teamRepo repos.Team,
) *Project {
	return &Project{
		logger:      logger,
		projectRepo: projectRepo,
		teamRepo:    teamRepo,
	}
}

func (p *Project) ByID(ctx context.Context, d db.DBTX, projectID int) (*db.Project, error) {
	defer newOTELSpan(ctx, "Project.ByID").End()

	project, err := p.projectRepo.ByID(ctx, d, projectID)
	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeNotFound, "project not found")
	}

	return project, nil
}

func (p *Project) ByName(ctx context.Context, d db.DBTX, name models.Project) (*db.Project, error) {
	defer newOTELSpan(ctx, "Project.ByName").End()

	project, err := p.projectRepo.ByName(ctx, d, name)
	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeNotFound, "project not found")
	}

	return project, nil
}

// obj1 : required existing database config
// update : config update object
// (we dont know any of those, just projectID)
// TODO call this for every project at app startup, therefore config is always up to date in the db. `update` will be empty in this first update.
// when a user updates the config in UI we use this same function but with empty [] of paths.
// merging logic is the same for both scenarios.
// we are not typing the update to save ourselves from manually adding a migration to change projects.board_config
// when _any_ field changes. we generate a new config the way it must be and merge with whatever was in db's board_config there at app startup.
// the endpoint to update it will be validated by openapi libs as usual.
func (p *Project) MergeConfigFields(ctx context.Context, d db.DBTX, projectName models.Project, update map[string]any) (*models.ProjectConfig, error) {
	project, err := p.projectRepo.ByName(ctx, d, projectName)
	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeNotFound, "project not found")
	}

	fieldsMap := make(map[string]map[string]any)

	fmt.Printf("project.BoardConfig: %v\n", project.BoardConfig)

	var workItem any
	// explicitly initialize what we want to allow an admin to edit in project config ui
	switch projectName {
	case models.ProjectDemo:
		workItem = &models.RestDemoWorkItemsResponse{DemoWorkItem: models.DbDemoWorkItem{}, Closed: pointers.New(time.Now())}
		// workItem = structs.InitializeFields(reflect.ValueOf(workItem), 1).Interface() // we want very specific fields to be editable in config so it doesn't clutter it
		fmt.Printf("workItem: %+v\n", workItem)
	case models.ProjectDemoTwo:
		fallthrough
	default:
		return nil, errors.New("not implemented")
	}
	pathKeys := structs.GetKeys(workItem, "")

	// index ProjectConfig.Fields by path for simpler logic
	for _, path := range pathKeys {
		fieldsMap[path] = defaultConfigField(path)
	}

	var boardConfigMap map[string]any
	fj, _ := json.Marshal(project.BoardConfig)
	json.Unmarshal(fj, &boardConfigMap)

	// update default config with current db config and merge updated config on top
	// merge with default config is necessary for project init,
	//  but merge with existing db config isn't really necessary.
	p.mergeFieldsMap(fieldsMap, boardConfigMap)
	p.mergeFieldsMap(fieldsMap, update)

	project.BoardConfig.Fields = make([]models.ProjectConfigField, 0, len(fieldsMap))
	for _, field := range fieldsMap {
		var fieldStruct models.ProjectConfigField

		fBlob, _ := json.Marshal(field)
		_ = json.Unmarshal(fBlob, &fieldStruct)

		project.BoardConfig.Fields = append(project.BoardConfig.Fields, fieldStruct)
	}

	return &project.BoardConfig, err
}

// defaultConfigField returns a map version of the ProjectConfig.Fields field.
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
func (p *Project) mergeFieldsMap(fieldsMap map[string]map[string]any, update map[string]any) {
	fieldsInterface, ok := update["fields"]
	if !ok {
		return
	}

	// map version of ProjectConfig.Fields field
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
			continue // unknown config field passed
		}

		if fm, ok := fieldsMap[path]; ok {
			newField := defaultConfigField(fm["path"].(string))

			for key, value := range fieldMap {
				if reflect.TypeOf(value) != reflect.TypeOf(newField[key]) {
					continue // our config was changed or a wrong type was provided
				}
				newField[key] = value
			}
			fieldsMap[path] = newField
		}
	}
}
