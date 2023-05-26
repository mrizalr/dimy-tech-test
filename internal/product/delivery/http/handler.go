package http

import "github.com/mrizalr/dimy-tech-test/internal/product"

type handler struct {
	usecase product.Usecase
}

func New(usecase product.Usecase) *handler {
	return &handler{
		usecase: usecase,
	}
}
