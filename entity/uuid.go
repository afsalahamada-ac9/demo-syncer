package entity

import "github.com/google/uuid"

// ID entity ID
type UUID = uuid.UUID

// NewID create a new entity UUID
func NewUUID() UUID {
	return UUID(uuid.New())
}

// StringToID convert a string to an entity UUID
func StringToUUID(s string) (UUID, error) {
	id, err := uuid.Parse(s)
	return UUID(id), err
}
