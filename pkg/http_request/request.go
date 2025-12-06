package http_request

import (
	"api_shop/pkg/pkgValidator"
	"api_shop/pkg/response"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindAndValidate[T any](c *gin.Context) *T {
	var req T

	// Bind (Will choose binder based on content-type)
	if err := c.ShouldBind(&req); err != nil {
		response.Error(c, http.StatusUnprocessableEntity, "خطا در داده ورودی", map[string]string{
			"body": "داده‌های ارسال‌شده معتبر نیستند",
		})
		return nil
	}

	// if req implements FormRequest, call Prepare() and Authorize()
	if fr, ok := any(&req).(pkgValidator.FormRequest); ok {
		// Prepare hook
		fr.Prepare()

		// Authorize
		if !fr.Authorize(c) {
			response.Error(c, http.StatusForbidden, "دسترسی ندارید", nil)
			return nil
		}
	}

	// Base validation using struct tags
	if err := pkgValidator.Validate.Struct(req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			response.ValidationError(c, ve)
			return nil
		}
		// other errors
		response.Error(c, http.StatusBadRequest, "خطای ولیدیشن", err.Error())
		return nil
	}

	// Apply extra rules from Rules() map: for each field => tag (validator.Var)
	if fr, ok := any(&req).(pkgValidator.FormRequest); ok {
		rules := fr.Rules()
		if len(rules) > 0 {
			// reflect to get field values
			val := reflect.ValueOf(req)
			// if req is a pointer, get elem
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}
			errorsMap := make(map[string][]string)
			for fieldName, tag := range rules {
				// find field by name (exported)
				f := val.FieldByName(fieldName)
				if !f.IsValid() {
					// try case-insensitive search (optional)
					// skip if not found
					continue
				}
				fieldInterface := f.Interface()

				if err := pkgValidator.Validate.Var(fieldInterface, tag); err != nil {
					// translate message(s)
					if ve, ok := err.(validator.ValidationErrors); ok {
						for _, e := range ve {
							// custom message override
							msg := e.Translate(pkgValidator.Translator)
							if m, ok := fr.Messages()[fieldName+"."+e.Tag()]; ok && m != "" {
								msg = m
							}
							errorsMap[fieldName] = append(errorsMap[fieldName], msg)
						}
					} else {
						// other error
						errorsMap[fieldName] = append(errorsMap[fieldName], err.Error())
					}
				}
			}
			if len(errorsMap) > 0 {
				// return combined errors (merge with ValidationError format)
				response.Error(c, http.StatusUnprocessableEntity, "خطا در داده‌های ورودی", errorsMap)
				return nil
			}
		}
	}

	// After hook
	if fr, ok := any(&req).(pkgValidator.FormRequest); ok {
		if err := fr.After(c); err != nil {
			// After can return error string
			response.Error(c, http.StatusUnprocessableEntity, "خطا پس از ولیدیشن", err.Error())
			return nil
		}
	}

	return &req
}
