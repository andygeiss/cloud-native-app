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

func Test_ModuleService_CreateModule_With_ValidConfig_Should_CreateAllFiles(t *testing.T) {
	// Arrange
	os.RemoveAll("testdata")
	svc := services.NewService(cfg)

	// Act
	err := svc.CreateModule()

	_, env := os.Stat(".env")
	_, gitignore := os.Stat(".gitignore")
	_, justfile := os.Stat(".justfile")
	_, index := os.Stat("cmd/server/assets/templates/index.tmpl")
	_, login := os.Stat("cmd/server/assets/templates/login.tmpl")
	_, styles := os.Stat("cmd/server/assets/static/styles.css")
	_, index_go := os.Stat("internal/app/adapters/inbound/ui/index.go")
	_, index_go_test := os.Stat("internal/app/adapters/inbound/ui/index_test.go")
	_, login_go := os.Stat("internal/app/adapters/inbound/ui/login.go")
	_, login_go_test := os.Stat("internal/app/adapters/inbound/ui/login_test.go")
	_, view_go := os.Stat("internal/app/adapters/inbound/ui/view.go")
	_, view_go_test := os.Stat("internal/app/adapters/inbound/ui/view_test.go")
	_, router := os.Stat("internal/app/adapters/inbound/router.go")
	_, router_test := os.Stat("internal/app/adapters/inbound/router_test.go")
	_, config := os.Stat("internal/app/config/config.go")
	_, maingo := os.Stat("cmd/server/main.go")
	_, gomod := os.Stat("go.mod")
	_, gosum := os.Stat("go.sum")
	_, git := os.Stat(".git")

	// Assert
	assert.That(t, "err must be nil", err, nil)
	assert.That(t, ".env must exists", env, nil)
	assert.That(t, ".gitignore must exists", gitignore, nil)
	assert.That(t, ".justfile must exists", justfile, nil)
	assert.That(t, "index.tmpl must exists", index, nil)
	assert.That(t, "login.tmpl must exists", login, nil)
	assert.That(t, "styles.css must exists", styles, nil)
	assert.That(t, "index.go must exists", index_go, nil)
	assert.That(t, "index_test.go must exists", index_go_test, nil)
	assert.That(t, "login.go must exists", login_go, nil)
	assert.That(t, "login_test.go must exists", login_go_test, nil)
	assert.That(t, "view.go must exists", view_go, nil)
	assert.That(t, "view_test.go must exists", view_go_test, nil)
	assert.That(t, "router.go must exists", router, nil)
	assert.That(t, "router_test.go must exists", router_test, nil)
	assert.That(t, "config.go must exists", config, nil)
	assert.That(t, "server/main.go must exists", maingo, nil)
	assert.That(t, "go.mod must exists", gomod, nil)
	assert.That(t, "go.sum must exists", gosum, nil)
	assert.That(t, ".git must exists", git, nil)
}
