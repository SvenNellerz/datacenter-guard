package models

type Template struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type TemplatesDatabase struct {
	Templates []Template `json:"templates"`
}

type GenerateRequest struct {
	Type       string            `json:"type" binding:"required"`
	TemplateID string            `json:"template_id"`
	Parameters map[string]string `json:"parameters"`
}

type GenerateResponse struct {
	Content  string `json:"content"`
	Filename string `json:"filename"`
}
