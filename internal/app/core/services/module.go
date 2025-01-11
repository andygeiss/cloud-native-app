package services

import (
	"sync"

	"github.com/andygeiss/cloud-native-app/internal/app/config"
	"github.com/andygeiss/cloud-native-app/internal/app/core/ports"
)

// ModuleService is the service for module interactions (e.g., CRUD operations).
type ModuleService struct {
	templatesPort ports.TemplatesPort
	mutex         sync.RWMutex
}

// NewService creates a new instance of ModuleService.
func NewService(cfg *config.Config) *ModuleService {
	return &ModuleService{
		templatesPort: cfg.TemplatesPort,
	}
}
