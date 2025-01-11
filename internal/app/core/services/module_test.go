package services_test

import (
	"embed"
	"os"
	"testing"

	"github.com/andygeiss/cloud-native-app/internal/app/adapters/outbound"
	"github.com/andygeiss/cloud-native-app/internal/app/config"
	"github.com/andygeiss/cloud-native-app/internal/app/core/services"
	"github.com/andygeiss/cloud-native-utils/assert"
)

//go:embed templates
var efs embed.FS

var cfg *config.Config

// init initializes the test by creating a new configuration.
func init() {
	cfg = &config.Config{
		Efs:       efs,
		Module:    "testdata",
		Templates: "templates/*",
	}
	cfg.TemplatesPort = outbound.NewTemplatesAdapter(cfg)
}

func TestModuleService_CreateModule(t *testing.T) {
	os.RemoveAll("testdata")

	// Create a new service with the configuration.
	service := services.NewService(cfg)

	// Create a new module.
	err := service.CreateModule()

	_, env := os.Stat(".env")
	_, gitignore := os.Stat(".gitignore")
	_, justfile := os.Stat(".justfile")
	_, dockerfile := os.Stat("Dockerfile")
	_, index := os.Stat("cmd/service/assets/index.html")
	_, handlers := os.Stat("internal/app/adapters/outbound/api/handlers.go")
	_, config := os.Stat("internal/app/config/config.go")
	_, genkey := os.Stat("cmd/genkey/main.go")
	_, maingo := os.Stat("cmd/service/main.go")
	_, gomod := os.Stat("go.mod")
	_, gosum := os.Stat("go.sum")
	_, git := os.Stat(".git")

	// Assert that the module was created successfully.
	assert.That(t, "err must be nil", err, nil)
	assert.That(t, ".env must exists", env, nil)
	assert.That(t, ".gitignore must exists", gitignore, nil)
	assert.That(t, ".justfile must exists", justfile, nil)
	assert.That(t, "Dockerfile must exists", dockerfile, nil)
	assert.That(t, "index.html must exists", index, nil)
	assert.That(t, "handlers.go must exists", handlers, nil)
	assert.That(t, "config.go must exists", config, nil)
	assert.That(t, "genkey/main.go must exists", genkey, nil)
	assert.That(t, "service/main.go must exists", maingo, nil)
	assert.That(t, "go.mod must exists", gomod, nil)
	assert.That(t, "go.sum must exists", gosum, nil)
	assert.That(t, ".git must exists", git, nil)
}
