package handler

import (
	"api_shop/htmx_sample/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
}

func Page(c *gin.Context) {

	u := []User{
		{Name: "ali"},
		{Name: "fati"},
		{Name: "reza"},
	}

	csrf := auth.GenerateCSRF(c)

	c.HTML(http.StatusOK, "pages", gin.H{
		"title": "user page",
		"csrf":  csrf,
		"users": u,
	})
}

func Store(c *gin.Context) {

	name := c.PostForm("name")

	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}
