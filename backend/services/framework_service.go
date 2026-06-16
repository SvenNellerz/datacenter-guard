package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/SvenNellerz/datacenter-guard/backend/models"
)

func LoadFrameworks() ([]models.Framework, error) {
	data, err := os.ReadFile(filepath.Join(dataDir(), "frameworks.json"))
	if err != nil {
		return nil, err
	}
	var db models.FrameworksDatabase
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}
	return db.Frameworks, nil
}

type ChecklistItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Checked     bool   `json:"checked"`
}

type ChecklistCategory struct {
	Name  string          `json:"name"`
	Items []ChecklistItem `json:"items"`
}

func GenerateChecklist(fw models.Framework) []ChecklistCategory {
	var result []ChecklistCategory
	cats := fw.Categories
	if cats == nil {
		cats = fw.Controls
	}
	for i, cat := range cats {
		var items []ChecklistItem
		for j, item := range cat.Items {
			switch v := item.(type) {
			case string:
				items = append(items, ChecklistItem{
					ID:    fmt.Sprintf("%s-%d-%d", fw.ID, i, j),
					Title: v,
				})
			case map[string]interface{}:
				title := ""
				desc := ""
				if t, ok := v["title"].(string); ok {
					title = t
				} else if t, ok := v["name"].(string); ok {
					title = t
				} else if t, ok := v["description"].(string); ok {
					title = t
				}
				if d, ok := v["description"].(string); ok {
					desc = d
				}
				items = append(items, ChecklistItem{
					ID:          fmt.Sprintf("%s-%d-%d", fw.ID, i, j),
					Title:       title,
					Description: desc,
				})
			}
		}
		result = append(result, ChecklistCategory{Name: cat.Name, Items: items})
	}
	return result
}
