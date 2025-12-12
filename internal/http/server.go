package http

import (
	"api_shop/config"
	"api_shop/internal/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	config.Load("http_request")

	r := gin.Default()

	routes.RegisterRoutes(r)

	log.Printf("http://localhost:%d", config.C.App.Port)
	log.Fatal(r.Run(fmt.Sprintf(":%d", config.C.App.Port)))

}
