package server

import (
	customerHttpHandler "github.com/mrizalr/dimy-tech-test/internal/customer/delivery/http"
	customerRepo "github.com/mrizalr/dimy-tech-test/internal/customer/repository/mysql"
	customerUcase "github.com/mrizalr/dimy-tech-test/internal/customer/usecase"
)

func (s *server) MapHandlers() {
	// Repository
	customerRepository := customerRepo.New(s.DB)

	// Usecase
	customerUsecase := customerUcase.New(customerRepository)

	// Handler
	customerHandler := customerHttpHandler.New(customerUsecase)

	// API Versioning
	api := s.App.Group("/api")
	v1 := api.Group("/v1")

	// Mapping Handlers
	customerGroup := v1.Group("/customers")
	customerHandler.MapRoutes(customerGroup)
}
