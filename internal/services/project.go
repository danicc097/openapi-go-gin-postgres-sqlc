package services

import (
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/models"
	"go.uber.org/zap"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Project struct {
	logger      *zap.Logger
	projectRepo repos.Project
	teamRepo    repos.Team
	authzsvc    *Authorization
}

// NewProject returns a new Project service.
func NewProject(logger *zap.Logger,
	projectRepo repos.Project,
	teamRepo repos.Team,
	authzsvc *Authorization,
) *Project {
	return &Project{
		logger:      logger,
		projectRepo: projectRepo,
		teamRepo:    teamRepo,
		authzsvc:    authzsvc,
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
func (p *Project) MergeConfigFields(obj1 models.ProjectConfig, obj2 any, pathKeys []string) models.ProjectConfig {
	fieldsMap := make(map[string]models.ProjectConfigField)
	tcaser := cases.Title(language.BrazilianPortuguese)
	if len(pathKeys) == 0 {
		for _, field := range obj1.Fields {
			fieldsMap[field.Path] = field
		}
	} else {
		for _, path := range pathKeys {
			fieldsMap[path] = models.ProjectConfigField{
				Path:          path,
				Name:          tcaser.String(path[strings.LastIndex(path, ".")+1:]),
				IsVisible:     true,
				IsEditable:    true,
				ShowCollapsed: true,
			}
		}

		p.mergeFieldsMap(fieldsMap, obj1)
	}

	p.mergeFieldsMap(fieldsMap, obj2)

	obj1.Fields = make([]models.ProjectConfigField, 0, len(fieldsMap))
	for _, field := range fieldsMap {
		obj1.Fields = append(obj1.Fields, field)
	}

	return obj1
}

func (p *Project) mergeFieldsMap(fieldsMap map[string]models.ProjectConfigField, obj interface{}) {
	objMap, ok := obj.(map[string]interface{})
	if !ok {
		return
	}

	fieldsInterface, ok := objMap["fields"]
	if !ok {
		return
	}

	fields, ok := fieldsInterface.([]interface{})
	if !ok {
		return
	}

	for _, fieldInterface := range fields {
		fieldMap, ok := fieldInterface.(map[string]interface{})
		if !ok {
			continue
		}

		path, ok := fieldMap["path"].(string)
		if !ok {
			continue
		}

		if fm, ok := fieldsMap[path]; ok {
			newField := models.ProjectConfigField{ // default
				Path:          fm.Path,
				IsEditable:    true,
				IsVisible:     true,
				ShowCollapsed: true,
				Name:          fm.Path,
			}
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
