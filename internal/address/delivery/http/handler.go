package http

import "github.com/mrizalr/dimy-tech-test/internal/address"

type handler struct {
	usecase address.Usecase
}

func New(usecase address.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}
