package services

type IdGenerator interface {
	GetRandomUUID() (string, error)
}
