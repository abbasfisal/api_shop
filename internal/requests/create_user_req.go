package requests

import "github.com/gin-gonic/gin"

type CreateUserRequest struct {
	Name  string `json:"name" validate:"required,min=3,max=20"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"required,min=0,max=99"`
}

func (r *CreateUserRequest) Prepare() {

}

func (r *CreateUserRequest) Authorize(c *gin.Context) bool {
	// مثال ساده: همیشه اجازه
	return true
}

func (r *CreateUserRequest) Rules() map[string]string {
	return map[string]string{
		"Name": "min=3,max=20",
		"Age":  "required,min=1",
	}
}

func (r *CreateUserRequest) Messages() map[string]string {
	return map[string]string{
		"Name.min":     "نام باید حداقل ۳ کاراکتر باشد",
		"Age.required": "سن باید وارد شود",
	}
}

func (r *CreateUserRequest) After(c *gin.Context) error {
	// برای مثال فرض کن میخوای چک کنی ایمیل یکتا باشه:
	// if emailExists(r.Email) { return errors.New("ایمیل قبلا ثبت شده") }
	return nil
}
