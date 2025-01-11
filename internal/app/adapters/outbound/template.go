package outbound

import (
	"bytes"
	"embed"

	"github.com/andygeiss/cloud-native-app/internal/app/config"
	"github.com/andygeiss/cloud-native-app/internal/app/core/ports"
	"github.com/andygeiss/cloud-native-utils/templating"
)

// TemplatesAdapter is the adapter for the templates port.
type TemplatesAdapter struct {
	efs    embed.FS
	engine *templating.Engine
}

// NewTemplatesAdapter creates a new instance of TemplatesAdapter.
func NewTemplatesAdapter(cfg *config.Config) ports.TemplatesPort {
	engine := templating.NewEngine(cfg.Efs)
	engine.Parse(cfg.Templates)
	return &TemplatesAdapter{
		efs:    cfg.Efs,
		engine: engine,
	}
}

// Get returns the rendered template as byte slice.
func (a *TemplatesAdapter) Get(name string, data any) (out []byte, err error) {
	// Create a new buffer.
	var buf bytes.Buffer

	// Render the template.
	if err := a.engine.Render(&buf, name, data); err != nil {
		return nil, err
	}

	// Return the buffer as byte slice.
	return buf.Bytes(), nil
}
