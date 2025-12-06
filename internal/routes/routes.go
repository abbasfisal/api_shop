package routes

import (
	"api_shop/internal/handlers"
	"api_shop/internal/handlers/user_handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")

	api.GET("/users", user_handler.Index)
	api.POST("/users", user_handler.Create)
	api.GET("/ping", handlers.PingHandler)
}
