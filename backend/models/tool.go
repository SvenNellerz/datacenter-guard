package models

type Tool struct {
	ID                      string   `json:"id"`
	Name                    string   `json:"name"`
	Vendor                  string   `json:"vendor"`
	Category                string   `json:"category"`
	Description             string   `json:"description"`
	Pricing                 string   `json:"pricing"`
	ComplexityScore         int      `json:"complexity_score"`
	KeyFeatures             []string `json:"key_features"`
	UseCases                []string `json:"use_cases"`
	IntegrationCapabilities []string `json:"integration_capabilities"`
	Pros                    []string `json:"pros"`
	Cons                    []string `json:"cons"`
	Certifications          []string `json:"certifications"`
	AlternativeTools        []string `json:"alternative_tools"`
	CostEstimate            string   `json:"cost_estimate"`
	DeploymentComplexity    string   `json:"deployment_complexity"`
}

type ToolsDatabase struct {
	Tools []Tool `json:"tools"`
}
