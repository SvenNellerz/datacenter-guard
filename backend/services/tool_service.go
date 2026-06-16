package services

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"

	"github.com/SvenNellerz/datacenter-guard/backend/models"
)

func dataDir() string {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, "..", "..", "frontend", "src", "assets", "data")
}

func LoadTools() ([]models.Tool, error) {
	data, err := os.ReadFile(filepath.Join(dataDir(), "tools.json"))
	if err != nil {
		return nil, err
	}
	var db models.ToolsDatabase
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}
	return db.Tools, nil
}
