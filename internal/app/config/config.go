package config

import (
	"embed"
	"path/filepath"

	"github.com/andygeiss/cloud-native-app/internal/app/core/ports"
)

// Config is the configuration for the application.
type Config struct {
	Efs           embed.FS            `json:"-"`
	Key           string              `json:"key"`
	Module        string              `json:"module"`
	Name          string              `json:"name"`
	Templates     string              `json:"templates"`
	TemplatesPort ports.TemplatesPort `json:"-"`
}

// GetName returns the project name derived from the module path.
func (a *Config) GetName() string {
	if a.Name != "" {
		return a.Name
	}
	return filepath.Base(a.Module)
}
