package lana


// Service is the default customerService interface
// implementation returned by customer.NewService.
type Service struct {
}

// NewService returns the default Service interface implementation.
func NewService() Service {
	return Service{}
}
