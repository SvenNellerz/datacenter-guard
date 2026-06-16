package models

type FrameworkItem struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type FrameworkCategory struct {
	Name  string        `json:"name"`
	Items []interface{} `json:"items"`
}

type Framework struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Categories  []FrameworkCategory `json:"categories,omitempty"`
	Controls    []FrameworkCategory `json:"controls,omitempty"`
}

type FrameworksDatabase struct {
	Frameworks []Framework `json:"frameworks"`
}
