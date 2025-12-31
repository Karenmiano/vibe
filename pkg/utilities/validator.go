package utilities

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)


func NewValidator() (*validator.Validate, ut.Translator) {
	validate := validator.New()

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate, trans
}

func TransformErrors(validateErrors validator.ValidationErrors, trans ut.Translator) map[string]string {
	errs := make(map[string]string)
	for _, fieldErr := range validateErrors {
		errs[fieldErr.Field()] = fieldErr.Translate(trans)
	}
	return errs
}