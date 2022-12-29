package services

import (
	"context"
	"encoding/json"
	"reflect"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/structs"
	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
func (p *Project) MergeConfigFields(ctx context.Context, d db.DBTX, projectID int, obj2 any) (*models.ProjectConfig, error) {
	project, err := p.projectRepo.ProjectByID(ctx, d, projectID)
	if err != nil {
		return nil, internal.NewErrorf(internal.ErrorCodeNotFound, "project not found")
	}

	fieldsMap := make(map[string]map[string]any)

	var obj1 models.ProjectConfig
	_ = json.Unmarshal(project.BoardConfig.Bytes, &obj1)

	pathKeys := structs.GetStructKeys(obj1, "")

	for _, path := range pathKeys {
		fieldsMap[path] = defaultConfigField(path)
	}
	p.mergeFieldsMap(fieldsMap, obj1)

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
	tcaser := cases.Title(language.English)
	f := models.ProjectConfigField{
		Path:          path,
		Name:          tcaser.String(path[strings.LastIndex(path, ".")+1:]),
		IsVisible:     true,
		IsEditable:    true,
		ShowCollapsed: true,
	}

	var jsonMap map[string]any

	fj, _ := json.Marshal(f)
	json.Unmarshal(fj, &jsonMap)

	return jsonMap
}

func (p *Project) mergeFieldsMap(fieldsMap map[string]map[string]any, obj any) {
	objMap, ok := obj.(map[string]any)
	if !ok {
		return
	}

	fieldsInterface, ok := objMap["fields"]
	if !ok {
		return
	}

	fields, ok := fieldsInterface.([]map[string]any)
	if !ok {
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
				newField[key] = value
			}
			fieldsMap[path] = newField
		}
	}
}
