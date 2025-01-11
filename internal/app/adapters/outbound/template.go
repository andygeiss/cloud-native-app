package outbound

import (
	"embed"

	"github.com/andygeiss/cloud-native-app/internal/app/config"
)

type TemplatesAdapter struct {
	efs embed.FS
}

func NewTemplatesAdapter(cfg *config.Config) *TemplatesAdapter {
	return &TemplatesAdapter{
		efs: cfg.Efs,
	}
}
