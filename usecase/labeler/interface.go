package labeler

import (
	"sudhagar/glad/entity"
)

// UseCase interface
type UseCase interface {
	SetLabel(tenantID entity.ID, contactID entity.ID, labelID entity.ID) error
	RemoveLabel(tenantID entity.ID, contactID entity.ID, labelID entity.ID) error
}
