package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"whatsappbot/src/server/routes"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	router *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "8080",
		router: gin.Default(),
	}
}

func (server *Server) Run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Encerrando servidor.")
			return
		default:
			router := routes.ConfigRoutes(server.router)
			log.Fatal(router.Run(":" + server.port))
		}
	}
}
