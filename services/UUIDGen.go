package services

import "github.com/google/uuid"

type UUIDGenerator interface {
	New() string
}

type UUIDValidator interface {
	// ValidateUUID indicates if a string is an UUID
	ValidateUUID(value string) bool
}

type UUIDGeneratorImpl struct{}

func (c UUIDGeneratorImpl) New() string {
	return uuid.New().String()
}

func (c UUIDGeneratorImpl) ValidateUUID(value string) bool {
	_, err := uuid.Parse(value)
	return err == nil
}
