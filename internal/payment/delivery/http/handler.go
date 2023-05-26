package http

import "github.com/mrizalr/dimy-tech-test/internal/payment"

type handler struct {
	usecase payment.Usecase
}

func New(usecase payment.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}
