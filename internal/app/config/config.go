package config

import (
	"embed"

	"github.com/andygeiss/cloud-native-app/internal/app/core/ports"
)

// Config is the configuration for the application.
type Config struct {
	Efs           embed.FS            `json:"-"`
	Key           string              `json:"key"`
	Module        string              `json:"module"`
	Templates     string              `json:"templates"`
	TemplatesPort ports.TemplatesPort `json:"-"`
}
