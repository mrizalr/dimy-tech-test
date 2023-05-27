package server

import (
	addressHttpHandler "github.com/mrizalr/dimy-tech-test/internal/address/delivery/http"
	addressRepo "github.com/mrizalr/dimy-tech-test/internal/address/repository/mysql"
	addressUcase "github.com/mrizalr/dimy-tech-test/internal/address/usecase"
	customerHttpHandler "github.com/mrizalr/dimy-tech-test/internal/customer/delivery/http"
	customerRepo "github.com/mrizalr/dimy-tech-test/internal/customer/repository/mysql"
	customerUcase "github.com/mrizalr/dimy-tech-test/internal/customer/usecase"
)

func (s *server) MapHandlers() {
	// Repository
	customerRepository := customerRepo.New(s.DB)
	addressRepository := addressRepo.New(s.DB)

	// Usecase
	customerUsecase := customerUcase.New(customerRepository)
	addressUsecase := addressUcase.New(addressRepository)

	// Handler
	customerHandler := customerHttpHandler.New(customerUsecase)
	addressHandler := addressHttpHandler.New(addressUsecase)

	// API Versioning
	api := s.App.Group("/api")
	v1 := api.Group("/v1")

	// Mapping Handlers
	customerGroup := v1.Group("/customers")
	customerHandler.MapRoutes(customerGroup)

	addressGroup := v1.Group("/addresses")
	addressHandler.MapRoutes(*addressGroup)
}
