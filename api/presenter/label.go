package presenter

import (
	"sudhagar/glad/entity"
)

// Label data - TenantID is returned in the HTTP header
// X-Messaging-TenantID
type Label struct {
	ID       entity.ID `json:"id"`
	TenantID entity.ID `json:"tenantId"`
	Name     string    `json:"name"`
	Color    uint32    `json:"color"`
}
