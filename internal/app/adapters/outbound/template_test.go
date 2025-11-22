package outbound_test

import (
	"embed"
	"testing"

	"github.com/andygeiss/cloud-native-app/internal/app/adapters/outbound"
	"github.com/andygeiss/cloud-native-app/internal/app/config"
	"github.com/andygeiss/cloud-native-utils/assert"
)

//go:embed testdata
var efs embed.FS

func TestTemplate_Get_OK(t *testing.T) {
	cfg := &config.Config{
		Efs:       efs,
		Templates: "testdata/*.tmpl",
	}
	a := outbound.NewTemplatesAdapter(cfg)

	out, err := a.Get("test", nil)
	assert.That(t, "err must be nil", err, nil)
	assert.That(t, "out must be OK", string(out), "OK")
}
