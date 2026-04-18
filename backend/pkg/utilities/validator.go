package utilities

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var rxUserName = regexp.MustCompile(`^[a-zA-Z0-9._]+$`)

// NewValidator creates a validator with error translations, for converting validation errors
// into readable messages.
func NewValidator() (*validator.Validate, ut.Translator) {
	validate := validator.New()

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(validate, trans)

	// Use JSON tag names instead of Go struct field names in validation errors
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Custom validator for username
	validate.RegisterValidation("username", func(fl validator.FieldLevel) bool {
		return rxUserName.MatchString(fl.Field().String())
	})

	// Custom translation for username
	validate.RegisterTranslation("username", trans,
		// Registration function, defines the translation
		func(ut ut.Translator) error {
			return ut.Add("username", "{0} must contain only letters, numbers, dots (.) and underscores (_)", true)
		},
		// Translation function - formats the message
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("username", fe.Field())
			return t
		},
	)

	return validate, trans
}

// Converts validator.ValidationErrors into a field-[messages...] mapping
func TransformErrors(validateErrors validator.ValidationErrors, trans ut.Translator) map[string]string {
	errs := make(map[string]string)
	for _, fieldErr := range validateErrors {
		errs[fieldErr.Field()] = fieldErr.Translate(trans)
	}
	return errs
}