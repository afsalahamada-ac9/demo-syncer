package presenter

import (
	"sudhagar/glad/entity"
)

// Template data - TenantID is returned in the HTTP header
// X-Messaging-TenantID
type Template struct {
	ID       entity.ID           `json:"id"`
	TenantID entity.ID           `json:"tenantId"`
	Name     string              `json:"name"`
	Type     entity.TemplateType `json:"type"`
	Content  string              `json:"content"`
}
