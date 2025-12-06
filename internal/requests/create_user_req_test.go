package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"api_shop/pkg/http_request"
	"api_shop/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	return r
}

func TestCreateUserRequest_ValidData(t *testing.T) {
	r := setupRouter()

	data := CreateUserRequest{
		Name:  "Ali",
		Email: "ali@example.com",
		Age:   25,
	}

	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.POST("/", func(c *gin.Context) {
		userReq := http_request.BindAndValidate[CreateUserRequest](c)
		assert.NotNil(t, userReq)
		assert.Equal(t, "Ali", userReq.Name)
		response.Success(c, userReq, "OK")
	})

	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateUserRequest_InvalidData(t *testing.T) {
	r := setupRouter()

	data := map[string]interface{}{
		"name":  "Al",            // کمتر از 3
		"email": "invalid-email", // اشتباه
		"age":   -5,              // اشتباه
	}

	body, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.POST("/", func(c *gin.Context) {
		userReq := http_request.BindAndValidate[CreateUserRequest](c)
		assert.Nil(t, userReq)
	})

	r.ServeHTTP(w, req)
	assert.Equal(t, 422, w.Code)
}
