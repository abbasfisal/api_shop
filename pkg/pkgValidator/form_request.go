package pkgValidator

import "github.com/gin-gonic/gin"

type FormRequest interface {
	Prepare()

	Rules() map[string]string

	Messages() map[string]string

	Authorize(c *gin.Context) bool

	After(c *gin.Context) error
}
