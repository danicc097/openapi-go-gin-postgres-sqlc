package services

import (
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
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
func (p *Project) MergeConfigFields(projectID string, obj2 any) models.ProjectConfig {
	fieldsMap := make(map[string]models.ProjectConfigField)

	// TODO get from db by project id
	obj1 := models.ProjectConfig{}
	pathKeys := structs.GetStructKeys(obj1, "")

	for _, path := range pathKeys {
		fieldsMap[path] = defaultConfigField(path)
	}
	p.mergeFieldsMap(fieldsMap, obj1)

	p.mergeFieldsMap(fieldsMap, obj2)

	obj1.Fields = make([]models.ProjectConfigField, 0, len(fieldsMap))
	for _, field := range fieldsMap {
		obj1.Fields = append(obj1.Fields, field)
	}

	return obj1
}

func defaultConfigField(path string) models.ProjectConfigField {
	tcaser := cases.Title(language.English)
	f := models.ProjectConfigField{
		Path:          path,
		Name:          tcaser.String(path[strings.LastIndex(path, ".")+1:]),
		IsVisible:     true,
		IsEditable:    true,
		ShowCollapsed: true,
	}

	return f
}

func (p *Project) mergeFieldsMap(fieldsMap map[string]models.ProjectConfigField, obj any) {
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
			newField := defaultConfigField(fm.Path)
			for key, value := range fieldMap {
				switch key {
				case "isEditable":
					if isEditable, ok := value.(bool); ok {
						newField.IsEditable = isEditable
					}
				case "showCollapsed":
					if showCollapsed, ok := value.(bool); ok {
						newField.ShowCollapsed = showCollapsed
					}
				case "isVisible":
					if isVisible, ok := value.(bool); ok {
						newField.IsVisible = isVisible
					}
				case "name":
					if name, ok := value.(string); ok {
						newField.Name = name
					}
				}

				fieldsMap[path] = newField
			}
		}
	}
}
