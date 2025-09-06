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

	// Get the gitignore template.
	env, err := a.templatesPort.Get("env", a.cfg)
	if err != nil {
		return err
	}

	// Write the gitignore to the module directory.
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

	os.MkdirAll("cmd/cli", 0755)

	cli, _ := a.templatesPort.Get("cli", a.cfg)
	if err := os.WriteFile("cmd/cli/main.go", cli, 0644); err != nil {
		return err
	}

	os.MkdirAll("cmd/genkey", 0755)

	// Get the genkey/main.go template.
	genkey, _ := a.templatesPort.Get("genkey.go", a.cfg)

	if err := os.WriteFile("cmd/genkey/main.go", genkey, 0644); err != nil {
		return err
	}

	os.MkdirAll("cmd/service/assets", 0755)

	// Get the index.html template.
	index := []byte(`{{ define "index" }}
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<script src="https://cdn.jsdelivr.net/npm/htmx.org@2.0.6/dist/htmx.min.js"></script>
	<link rel="stylesheet" href="/assets/styles.css">
	<title> Messaging Demo </title>
</head>
<body>
	<button hx-post="/api/producer" hx-target="#response"> Send Message </button>
	<div id="response"></div>
</body>
</html>
{{ end }}
`)

	// Write the index.html to the module directory.
	if err := os.WriteFile("cmd/service/assets/index.html", index, 0644); err != nil {
		return err
	}

	// Write the styles.css.
	styles, _ := a.templatesPort.Get("styles", a.cfg)

	if err := os.WriteFile("cmd/service/assets/styles.css", styles, 0644); err != nil {
		return err
	}

	// Get the main.go template.
	maingo, err := a.templatesPort.Get("main.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the main.go to the module directory.
	if err := os.WriteFile("cmd/service/main.go", maingo, 0644); err != nil {
		return err
	}

	os.MkdirAll("internal/app/adapters/api", 0755)

	// Get the handlers.go template.
	handlers, err := a.templatesPort.Get("handlers.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the handlers.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/api/handlers.go", handlers, 0644); err != nil {
		return err
	}

	middleware, err := a.templatesPort.Get("middleware.go", a.cfg)
	if err != nil {
		return err
	}

	if err := os.WriteFile("internal/app/adapters/api/middleware.go", middleware, 0644); err != nil {
		return err
	}

	// Get the router.go template.
	router, err := a.templatesPort.Get("router.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the router.go to the module directory.
	if err := os.WriteFile("internal/app/adapters/api/router.go", router, 0644); err != nil {
		return err
	}

	os.MkdirAll("internal/app/config", 0755)

	// Get the handlers.go template.
	config, err := a.templatesPort.Get("config.go", a.cfg)
	if err != nil {
		return err
	}

	// Write the handlers.go to the module directory.
	if err := os.WriteFile("internal/app/config/config.go", config, 0644); err != nil {
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
