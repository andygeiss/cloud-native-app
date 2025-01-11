package config

import (
	"embed"

	"github.com/andygeiss/cloud-native-app/internal/app/core/ports"
)

// Config is the configuration for the application.
type Config struct {
	Efs           embed.FS            `json:"-"`
	TemplatesPort ports.TemplatesPort `json:"-"`
	Name          string              `json:"name"`
}
