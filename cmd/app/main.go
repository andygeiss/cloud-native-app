package main

import (
	"embed"
	"log"
	"os"
	"path/filepath"

	"github.com/andygeiss/cloud-native-app/internal/app/config"
)

//go:embed resources
var efs embed.FS

var version string

func main() {
	cfg := &config.Config{
		Efs:  efs,
		Name: filepath.Base(os.Getenv("PWD")),
	}
	log.Printf("%s (%s)", cfg.Name, version)
}
