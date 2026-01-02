package services

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/andygeiss/cloud-native-app/internal/app/config"
	"github.com/andygeiss/cloud-native-app/internal/app/core/ports"
)

// ModuleService is the service for module interactions (e.g., CRUD operations).
type ModuleService struct {
	cfg           *config.Config
	templatesPort ports.TemplatesPort
}

// NewService creates a new instance of ModuleService.
func NewService(cfg *config.Config) *ModuleService {
	return &ModuleService{
		cfg:           cfg,
		templatesPort: cfg.TemplatesPort,
	}
}

// CreateModule creates a new module.
func (a *ModuleService) CreateModule() error {
	// Change the working directory to the module directory.
	basename := filepath.Base(a.cfg.Module)
	_ = os.MkdirAll(basename, 0755)
	if err := os.Chdir(basename); err != nil {
		return err
	}

	// Get the env template.
	env, err := a.templatesPort.Get("env", a.cfg)
	if err != nil {
		return err
	}

	// Write the env to the module directory.
	if err := os.WriteFile(".env", env, 0644); err != nil {
		return err
	}

	// Get the gitignore template.
	gitignore, err := a.templatesPort.Get("gitignore", a.cfg)
	if err != nil {
		return err
	}

	// Write the gitignore to the module directory.
	if err := os.WriteFile(".gitignore", gitignore, 0644); err != nil {
		return err
	}

	// Get the justfile template.
	justfile, err := a.templatesPort.Get("justfile", a.cfg)
	if err != nil {
		return err
	}

	// Write the justfile to the module directory.
	if err := os.WriteFile(".justfile", justfile, 0644); err != nil {
		return err
	}

	// Get the readme template.
	readme, err := a.templatesPort.Get("readme.md", a.cfg)
	if err != nil {
		return err
	}

	// Write the README.md to the module directory.
	if err := os.WriteFile("README.md", readme, 0644); err != nil {
		return err
	}

	os.MkdirAll("cmd/server/assets/static", 0755)
	os.MkdirAll("cmd/server/assets/templates", 0755)

	// Get the index.tmpl HTML template.
	indexTmpl, err := a.templatesPort.Get("index.tmpl", a.cfg)
	if err != nil {
		return err
	}

	// Write the index.tmpl to the module directory.
	if err := os.WriteFile("cmd/server/assets/templates/index.tmpl", indexTmpl, 0644); err != nil {
		return err
	}

	// Get the login.tmpl HTML template.
	loginTmpl, err := a.templatesPort.Get("login.tmpl", a.cfg)
	if err != nil {
		return err
	}

	// Write the login.tmpl to the module directory.
	if err := os.WriteFile("cmd/server/assets/templates/login.tmpl", loginTmpl, 0644); err != nil {
		return err
	}

	// Get the base.css template.
	baseCss, err := a.templatesPort.Get("base.css", a.cfg)
	if err != nil {
		return err
	}

	// Write the base.css to the module directory.
	if err := os.WriteFile("cmd/server/assets/static/base.css", baseCss, 0644); err != nil {
		return err
	}

	// Get the theme.css template.
	themeCss, err := a.templatesPort.Get("theme.css", a.cfg)
	if err != nil {
		return err
	}

	// Write the theme.css to the module directory.
	if err := os.WriteFile("cmd/server/assets/static/theme.css", themeCss, 0644); err != nil {
		return err
	}

	// Get the styles.css template.
	stylesCss, err := a.templatesPort.Get("styles.css", a.cfg)
	if err != nil {
		return err
	}

	// Write the styles.css to the module directory.
	if err := os.WriteFile("cmd/server/assets/static/styles.css", stylesCss, 0644); err != nil {
		return err
	}

	// Get the main.go template.
	maingo, err := a.templatesPort.Get("main.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the main.go to the server directory.
	if err := os.WriteFile("cmd/server/main.go", maingo, 0644); err != nil {
		return err
	}

	os.MkdirAll("internal/app", 0755)
	os.MkdirAll("internal/app/adapters/inbound/ui", 0755)

	// Get the errors.go template.
	errorsGo, err := a.templatesPort.Get("errors.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the errors.go to the module directory.
	if err := os.WriteFile("internal/app/errors.go", errorsGo, 0644); err != nil {
		return err
	}

	// Get the index.go template.
	indexGo, err := a.templatesPort.Get("index.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the index.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/inbound/ui/index.go", indexGo, 0644); err != nil {
		return err
	}

	// Get the index_test.go template.
	indexTestGo, err := a.templatesPort.Get("index_test.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the index_test.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/inbound/ui/index_test.go", indexTestGo, 0644); err != nil {
		return err
	}

	// Get the login.go template.
	loginGo, err := a.templatesPort.Get("login.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the login.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/inbound/ui/login.go", loginGo, 0644); err != nil {
		return err
	}

	// Get the login_test.go template.
	loginTestGo, err := a.templatesPort.Get("login_test.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the login_test.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/inbound/ui/login_test.go", loginTestGo, 0644); err != nil {
		return err
	}

	// Get the middleware.go template.
	middlewareGo, err := a.templatesPort.Get("middleware.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the middleware.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/inbound/middleware.go", middlewareGo, 0644); err != nil {
		return err
	}

	// Get the view.go template.
	viewGo, err := a.templatesPort.Get("view.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the view.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/inbound/ui/view.go", viewGo, 0644); err != nil {
		return err
	}

	// Get the view_test.go template.
	viewTestGo, err := a.templatesPort.Get("view_test.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the view_test.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/inbound/ui/view_test.go", viewTestGo, 0644); err != nil {
		return err
	}

	// Get the router.go template.
	routerGo, err := a.templatesPort.Get("router.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the router.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/inbound/router.go", routerGo, 0644); err != nil {
		return err
	}

	// Get the router_test.go template.
	routerTestGo, err := a.templatesPort.Get("router_test.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the router_test.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/inbound/router_test.go", routerTestGo, 0644); err != nil {
		return err
	}

	os.MkdirAll("internal/app/adapters/inbound/testdata", 0755)
	os.MkdirAll("internal/app/adapters/inbound/ui/testdata", 0755)

	// Write the index.tmpl to the testdata directories.
	if err := os.WriteFile("internal/app/adapters/inbound/testdata/index.tmpl", indexTmpl, 0644); err != nil {
		return err
	}
	if err := os.WriteFile("internal/app/adapters/inbound/ui/testdata/index.tmpl", indexTmpl, 0644); err != nil {
		return err
	}

	// Write the login.tmpl to the testdata directories.
	if err := os.WriteFile("internal/app/adapters/inbound/testdata/login.tmpl", loginTmpl, 0644); err != nil {
		return err
	}
	if err := os.WriteFile("internal/app/adapters/inbound/ui/testdata/login.tmpl", loginTmpl, 0644); err != nil {
		return err
	}

	os.MkdirAll("internal/app/config", 0755)

	// Get the config.go template.
	configGo, err := a.templatesPort.Get("config.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the config.go to the module directory.
	if err := os.WriteFile("internal/app/config/config.go", configGo, 0644); err != nil {
		return err
	}

	// Initialize the module.
	exec.Command("go", "mod", "init", a.cfg.Module).Run()
	exec.Command("go", "mod", "tidy").Run()
	exec.Command("git", "init").Run()
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", "feat: init repo").Run()

	return nil
}
