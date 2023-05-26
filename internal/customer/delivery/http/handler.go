package http

import "github.com/mrizalr/dimy-tech-test/internal/customer"

type handler struct {
	usecase customer.Usecase
}

func New(usecase customer.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}
