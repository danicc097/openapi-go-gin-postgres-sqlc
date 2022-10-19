package services

// SomeService is a dummy service to showcase calling services from others.
type SomeService struct{}

// NewSomeService returns a new SomeService service.
func NewSomeService() *SomeService {
	return &SomeService{}
}
