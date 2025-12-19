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

func Test_TemplatesAdapter_Get_With_ValidTemplate_Should_ReturnContent(t *testing.T) {
	// Arrange
	cfg := &config.Config{
		Efs:       efs,
		Templates: "testdata/*.tmpl",
	}
	a := outbound.NewTemplatesAdapter(cfg)

	// Act
	out, err := a.Get("test", nil)

	// Assert
	assert.That(t, "err must be nil", err, nil)
	assert.That(t, "out must be OK", string(out), "OK")
}
