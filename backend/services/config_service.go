package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/SvenNellerz/datacenter-guard/backend/models"
)

func LoadTemplates() ([]models.Template, error) {
	data, err := os.ReadFile(filepath.Join(dataDir(), "templates.json"))
	if err != nil {
		return nil, err
	}
	var db models.TemplatesDatabase
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}
	return db.Templates, nil
}

func GenerateConfig(req models.GenerateRequest) (models.GenerateResponse, error) {
	templates, err := LoadTemplates()
	if err != nil {
		return models.GenerateResponse{}, err
	}

	var content string
	filename := fmt.Sprintf("config-%s.tf", req.Type)

	if req.TemplateID != "" {
		for _, t := range templates {
			if t.ID == req.TemplateID {
				content = applyParameters(t.Content, req.Parameters)
				filename = fmt.Sprintf("%s.%s", t.ID, fileExtension(t.Type))
				break
			}
		}
	}

	if content == "" {
		content = defaultConfig(req.Type)
	}

	return models.GenerateResponse{Content: content, Filename: filename}, nil
}

func applyParameters(content string, params map[string]string) string {
	for k, v := range params {
		content = strings.ReplaceAll(content, fmt.Sprintf("{{%s}}", k), v)
	}
	return content
}

func fileExtension(configType string) string {
	switch strings.ToLower(configType) {
	case "terraform":
		return "tf"
	case "ansible":
		return "yml"
	default:
		return "txt"
	}
}

func defaultConfig(configType string) string {
	switch strings.ToLower(configType) {
	case "terraform":
		return `terraform {
  required_version = ">= 1.5.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

variable "aws_region" {
  default = "us-east-1"
}
`
	case "ansible":
		return `---
- name: Security Hardening
  hosts: all
  become: yes
  tasks:
    - name: Update packages
      package:
        name: '*'
        state: latest
`
	default:
		return "# Configuration template"
	}
}

func LoadArchitectures() ([]map[string]interface{}, error) {
	data, err := os.ReadFile(filepath.Join(dataDir(), "architectures.json"))
	if err != nil {
		return nil, err
	}
	var db struct {
		Templates []map[string]interface{} `json:"templates"`
	}
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}
	return db.Templates, nil
}

func LoadPlaybooks() ([]map[string]interface{}, error) {
	data, err := os.ReadFile(filepath.Join(dataDir(), "playbooks.json"))
	if err != nil {
		return nil, err
	}
	var db struct {
		Playbooks []map[string]interface{} `json:"playbooks"`
	}
	if err := json.Unmarshal(data, &db); err != nil {
		return nil, err
	}
	return db.Playbooks, nil
}
