package services

// IdGenerator offers UUID generation
type IdGenerator interface {
	GetRandomUUID() (string, error)
}
