package main

import (
	"api_shop/htmx_sample/internal/auth"
	"api_shop/htmx_sample/internal/handler"
	"api_shop/htmx_sample/internal/tmpl"
	"fmt"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//
	r.Use(auth.SessionMiddleware())

	//

	//template
	tmplEngine := tmpl.TmplEngine()
	for _, tmpl := range tmplEngine.Templates() {
		fmt.Println("loaded:", tmpl.Name())
	}
	r.SetHTMLTemplate(tmplEngine)

	//
	web := r.Group("web/")
	web.GET("/", handler.Page)
	web.POST("/create", auth.VerifyCSRF(), handler.Store)

	//api

	//server
	fmt.Println("http://localhost:8087")
	log.Fatal(r.Run(":8087"))

}
