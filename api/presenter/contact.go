package presenter

import (
	"sudhagar/glad/entity"
)

// Contact data - TenantID is returned in the HTTP header
// X-Messaging-TenantID
type Contact struct {
	ID        entity.ID `json:"id"`
	TenantID  entity.ID `json:"tenantId"`
	AccountID entity.ID `json:"accountId"`
	Handle    string    `json:"handle"`
	Name      string    `json:"name"`
}
