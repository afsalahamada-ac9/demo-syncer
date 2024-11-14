package presenter

import (
	"sudhagar/glad/entity"
)

// Tenant data
type Tenant struct {
	ID       entity.ID `json:"id"`
	Username string    `json:"username"`
	// Do not return password
	// AuthToken is returned at login
	AuthToken string `json:"token,omitempty"`
}
