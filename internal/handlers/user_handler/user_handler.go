package user_handler

import (
	"api_shop/config"
	"api_shop/internal/requests"
	"api_shop/pkg/http_request"

	"api_shop/pkg/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	users := []map[string]any{
		{
			"id":   1,
			"name": "ali",
			"age":  10,
		},
		{
			"id":   2,
			"name": "reza",
			"age":  32,
		},
		{
			"id":   122,
			"name": "mohammad",
			"age":  12,
		},
		{
			"id":   1,
			"name": "fat",
			"age":  10,
		},
		{
			"id":   1,
			"name": "sdsdf",
			"age":  10,
		},
		{
			"id":   1,
			"name": "sssss",
			"age":  10,
		},
		{
			"id":   2,
			"name": "aaaaa",
			"age":  32,
		},
		{
			"id":   122,
			"name": "bvbvbv",
			"age":  12,
		},
		{
			"id":   1,
			"name": "we3rwez",
			"age":  10,
		},
		{
			"id":   1,
			"name": "asdfasdfqw233",
			"age":  10,
		},
	}

	page := 1
	perPage := 5

	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if pp := c.Query("per_page"); pp != "" {
		fmt.Sscanf(pp, "%d", &perPage)
	}

	response.Paginate(c, users, page, perPage, fmt.Sprintf("http://localhost:%d/api/v1/users/", config.C.App.Port))
	return
}

func Create(c *gin.Context) {
	req := http_request.BindAndValidate[requests.CreateUserRequest](c)
	if req == nil {
		return
	}

	response.Success(c, req, "کاربر با موفقیت ایجاد شد")
}
