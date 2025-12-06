package pkgValidator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
)

var (
	Validate   *validator.Validate
	Translator ut.Translator
)

func init() {
	if false {
		Validate = validator.New()

		faLocale := fa.New()
		uni := ut.New(faLocale, faLocale)
		var found bool
		Translator, found = uni.GetTranslator("fa")
		if !found {
			panic("translator فارسی یافت نشد")
		}

		if err := fa_translations.RegisterDefaultTranslations(Validate, Translator); err != nil {
			panic(err)
		}

	}

	if true {
		Validate = validator.New()

		enLocale := en.New()
		uni := ut.New(enLocale, enLocale)
		var found bool
		Translator, found = uni.GetTranslator("en")
		if !found {
			panic("translator not found")
		}

		if err := en_translations.RegisterDefaultTranslations(Validate, Translator); err != nil {
			panic(err)
		}

	}

}

func RegisterValidation(tag string, fn validator.Func) error {
	return Validate.RegisterValidation(tag, fn)
}
