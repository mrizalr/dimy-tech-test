package server

import (
	addressHttpHandler "github.com/mrizalr/dimy-tech-test/internal/address/delivery/http"
	addressRepo "github.com/mrizalr/dimy-tech-test/internal/address/repository/mysql"
	addressUcase "github.com/mrizalr/dimy-tech-test/internal/address/usecase"
	customerHttpHandler "github.com/mrizalr/dimy-tech-test/internal/customer/delivery/http"
	customerRepo "github.com/mrizalr/dimy-tech-test/internal/customer/repository/mysql"
	customerUcase "github.com/mrizalr/dimy-tech-test/internal/customer/usecase"
	orderHttpHandler "github.com/mrizalr/dimy-tech-test/internal/order/delivery/http"
	orderRepo "github.com/mrizalr/dimy-tech-test/internal/order/repository/mysql"
	orderUcase "github.com/mrizalr/dimy-tech-test/internal/order/usecase"
	paymentHttpHandler "github.com/mrizalr/dimy-tech-test/internal/payment/delivery/http"
	paymentRepo "github.com/mrizalr/dimy-tech-test/internal/payment/repository/mysql"
	paymentUcase "github.com/mrizalr/dimy-tech-test/internal/payment/usecase"
	productHttpHandler "github.com/mrizalr/dimy-tech-test/internal/product/delivery/http"
	productRepo "github.com/mrizalr/dimy-tech-test/internal/product/repository/mysql"
	productUcase "github.com/mrizalr/dimy-tech-test/internal/product/usecase"
)

func (s *server) MapHandlers() {
	// Repository
	customerRepository := customerRepo.New(s.DB)
	addressRepository := addressRepo.New(s.DB)
	productRepository := productRepo.New(s.DB)
	paymentRepository := paymentRepo.New(s.DB)
	orderRepository := orderRepo.New(s.DB)

	// Usecase
	customerUsecase := customerUcase.New(customerRepository)
	addressUsecase := addressUcase.New(addressRepository)
	productUsecase := productUcase.New(productRepository)
	paymentUsecase := paymentUcase.New(paymentRepository)
	orderUsecase := orderUcase.New(orderRepository, productRepository)

	// Handler
	customerHandler := customerHttpHandler.New(customerUsecase)
	addressHandler := addressHttpHandler.New(addressUsecase)
	productHandler := productHttpHandler.New(productUsecase)
	paymentHandler := paymentHttpHandler.New(paymentUsecase)
	orderHandler := orderHttpHandler.New(orderUsecase)

	// API Versioning
	api := s.App.Group("/api")
	v1 := api.Group("/v1")

	// Mapping Handlers
	customerGroup := v1.Group("/customers")
	customerHandler.MapRoutes(customerGroup)

	addressGroup := v1.Group("/addresses")
	addressHandler.MapRoutes(addressGroup)

	productGroup := v1.Group("/products")
	productHandler.MapRoutes(productGroup)

	paymentGroup := v1.Group("/payment-methods")
	paymentHandler.MapRoutes(paymentGroup)

	orderGroup := v1.Group("/orders")
	orderHandler.MapRoutes(orderGroup)
}
