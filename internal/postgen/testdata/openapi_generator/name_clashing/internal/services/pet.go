package services

type Pet struct{}

// NewPet returns a new Pet service.
func NewPet() *Pet {
	return &Pet{}
}
