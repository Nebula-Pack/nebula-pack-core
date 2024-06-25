package models

import (
	"fmt"
)

type NebulaConfig struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Description     string            `json:"description"`
	Main            string            `json:"main"`
	Scripts         map[string]string `json:"scripts"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Author          string            `json:"author"`
	License         string            `json:"license"`
}

// Process simulates processing of Nebula config (e.g., deducting dependencies)
func (c *NebulaConfig) Process() error {
	fmt.Println("Processing Nebula Config...")
	// Here you would deduct dependencies, manage scripts, etc.
	return nil
}
