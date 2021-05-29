package infrastructure

import (
	uuid "github.com/nu7hatch/gouuid"
)

type IdGenerator struct{}

func NewIdGenerator() *IdGenerator {
	return &IdGenerator{}
}

func (i *IdGenerator) GetRandomUUID() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}
