/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"sudhagar/glad/entity"
)

// Course data - TenantID is returned in the HTTP header
// X-GLAD-TenantID
type Course struct {
	ID           entity.ID            `json:"id"`
	ExtID        *string              `json:"extId,omitempty"`
	CenterID     *entity.ID           `json:"centerId,omitempty"`
	Name         *string              `json:"name,omitempty"`
	Notes        *string              `json:"notes,omitempty"`
	Timezone     *string              `json:"timezone,omitempty"`
	Address      *Address             `json:"address,omitempty"`
	Status       *entity.CourseStatus `json:"status,omitempty"`
	Mode         *entity.CourseMode   `json:"mode,omitempty"`
	MaxAttendees *int32               `json:"maxAttendees,omitempty"`
	NumAttendees *int32               `json:"numAttendees,omitempty"`
}
