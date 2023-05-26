package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type server struct {
	DB  *gorm.DB
	App *gin.Engine
}

func New(DB *gorm.DB) *server {
	return &server{
		DB:  DB,
		App: gin.Default(),
	}
}

func (s *server) Run() error {
	s.MapHandlers()

	port := os.Getenv("SERVER_PORT")

	log.Printf("Server is running on port %s", port)
	return s.App.Run(fmt.Sprintf(":%s", port))

}
