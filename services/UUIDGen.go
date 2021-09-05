package services

import "github.com/google/uuid"

type UUIDGenerator interface {
	New() string
}

type UUIDGeneratorImpl struct{}

func (c UUIDGeneratorImpl) New() string {
	return uuid.New().String()
}
