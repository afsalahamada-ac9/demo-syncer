/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"sudhagar/glad/entity"
)

// Account data - TenantID is returned in the HTTP header
// X-Messaging-TenantID
type Account struct {
	ID       entity.ID            `json:"id"`
	TenantID entity.ID            `json:"tenantId"`
	Username string               `json:"username"`
	Type     entity.AccountType   `json:"type"`
	Data     string               `json:"data,omitempty"`
	Status   entity.AccountStatus `json:"status"`
}
