package models

type ProjectConfig struct {
	Header []string             `json:"header" required:"true"`
	Fields []ProjectConfigField `json:"fields" required:"true"`
}

type ProjectConfigField struct {
	IsEditable    bool   `json:"isEditable" required:"true"`
	ShowCollapsed bool   `json:"showCollapsed" required:"true"`
	IsVisible     bool   `json:"isVisible" required:"true"`
	Path          string `json:"path" required:"true"`
	Name          string `json:"name" required:"true"`
}
