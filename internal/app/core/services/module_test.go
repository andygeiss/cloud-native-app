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
	_, index := os.Stat("cmd/service/static/index.tmpl")
	_, login := os.Stat("cmd/service/static/login.tmpl")
	_, styles := os.Stat("cmd/service/static/styles.css")
	_, index_go := os.Stat("internal/app/adapters/ingres/ui/index.go")
	_, index_go_test := os.Stat("internal/app/adapters/ingres/ui/index_test.go")
	_, login_go := os.Stat("internal/app/adapters/ingres/ui/login.go")
	_, login_go_test := os.Stat("internal/app/adapters/ingres/ui/login_test.go")
	_, view_go := os.Stat("internal/app/adapters/ingres/ui/view.go")
	_, view_go_test := os.Stat("internal/app/adapters/ingres/ui/view_test.go")
	_, router := os.Stat("internal/app/adapters/ingres/router.go")
	_, router_test := os.Stat("internal/app/adapters/ingres/router_test.go")
	_, config := os.Stat("internal/app/config/config.go")
	_, cli := os.Stat("cmd/cli/main.go")
	_, maingo := os.Stat("cmd/service/main.go")
	_, gomod := os.Stat("go.mod")
	_, gosum := os.Stat("go.sum")
	_, git := os.Stat(".git")

	// Assert that the module was created successfully.
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
	assert.That(t, "cli/main.go must exists", cli, nil)
	assert.That(t, "service/main.go must exists", maingo, nil)
	assert.That(t, "go.mod must exists", gomod, nil)
	assert.That(t, "go.sum must exists", gosum, nil)
	assert.That(t, ".git must exists", git, nil)
}
