package main

import (
	"embed"
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/andygeiss/cloud-native-app/internal/app/adapters/outbound"
	"github.com/andygeiss/cloud-native-app/internal/app/config"
	"github.com/andygeiss/cloud-native-app/internal/app/core/services"
	"github.com/andygeiss/cloud-native-utils/security"
)

//go:embed templates
var efs embed.FS

func main() {
	// Check if the module name is provided.
	if len(os.Args) != 2 {
		panic("module name is missing")
	}
	module := os.Args[1]

	// Create a configuration.
	key := security.GenerateKey()
	cfg := &config.Config{
		Efs:       efs,
		Key:       hex.EncodeToString(key[:]),
		Module:    module,
		Name:      filepath.Base(module),
		Templates: "templates/*",
	}
	cfg.TemplatesPort = outbound.NewTemplatesAdapter(cfg)

	// Create a new service.
	svc := services.NewService(cfg)

	// Create a new module.
	if err := svc.CreateModule(); err != nil {
		panic(err)
	}
}
