package response

import (
	"api_shop/pkg/pkgValidator"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Total          int    `json:"total,omitempty"`
	PerPage        int    `json:"per_page,omitempty"`
	CurrentPage    int    `json:"current_page,omitempty"`
	LastPage       int    `json:"last_page,omitempty"`
	CurrentPageURL string `json:"current_page_url,omitempty"`
	FirstPageURL   string `json:"first_page_url,omitempty"`
	LastPageURL    string `json:"last_page_url,omitempty"`
	NextPageURL    string `json:"next_page_url,omitempty"`
	PrevPageURL    string `json:"prev_page_url,omitempty"`
	Path           string `json:"path,omitempty"`
	From           int    `json:"from,omitempty"`
	To             int    `json:"to,omitempty"`
}

type ErrorItem struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(200, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, data interface{}, message string) {
	c.JSON(201, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string, errors interface{}) {
	c.JSON(code, APIResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}
func InvalidInputError(c *gin.Context) {
	Error(c, 422, "خطا در داده ورودی", map[string]string{
		"body": "فیلدها نامعتبر هستند",
	})
}

func Paginate(c *gin.Context, items []map[string]interface{}, page, perPage int, path string) {
	total := len(items)
	lastPage := total / perPage
	if total%perPage != 0 {
		lastPage++
	}

	start := (page - 1) * perPage
	end := start + perPage
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	pageItems := items[start:end]

	currentPageURL := fmt.Sprintf("%s?page=%d", path, page)
	firstPageURL := fmt.Sprintf("%s?page=1", path)
	lastPageURL := fmt.Sprintf("%s?page=%d", path, lastPage)
	nextPageURL := ""
	if page < lastPage {
		nextPageURL = fmt.Sprintf("%s?page=%d", path, page+1)
	}
	prevPageURL := ""
	if page > 1 {
		prevPageURL = fmt.Sprintf("%s?page=%d", path, page-1)
	}

	from := start + 1
	to := end

	meta := Meta{
		Total:          total,
		PerPage:        perPage,
		CurrentPage:    page,
		LastPage:       lastPage,
		CurrentPageURL: currentPageURL,
		FirstPageURL:   firstPageURL,
		LastPageURL:    lastPageURL,
		NextPageURL:    nextPageURL,
		PrevPageURL:    prevPageURL,
		Path:           path,
		From:           from,
		To:             to,
	}

	c.JSON(200, APIResponse{
		Success: true,
		Data:    pageItems,
		Meta:    &meta,
		Message: "success",
	})

}

func ValidationError(c *gin.Context, errs validator.ValidationErrors) {

	errors := make(map[string][]string)
	for _, err := range errs {
		field := err.Field()
		msg := err.Translate(pkgValidator.Translator)
		errors[field] = append(errors[field], msg)
	}

	c.JSON(http.StatusUnprocessableEntity, APIResponse{
		Success: false,
		Message: "خطا در داده‌های ورودی",
		Errors:  errors,
	})
}
