package utils

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func Validate(data interface{}) []string {
	en := en.New()
	uni := ut.New(en, en)

	trans, _ := uni.GetTranslator("en")
	validate := validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "" {
			return ""
		}

		return name
	})

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} is not email format", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())

		return t
	})

	validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "{0} must have min {1} character", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Param())

		return t
	})

	validate.RegisterTranslation("password", trans, func(ut ut.Translator) error {
		return ut.Add("password", "{0} must contain minmun 8 character, 1 capital character, 1 number, and 1 symbol", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("password", fe.Field())

		return t
	})

	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		if len(password) < 9 {
			return false
		}

		var (
			hasUppercase bool
			hasNumber    bool
			hasSymbol    bool
		)

		for _, char := range password {
			if unicode.IsUpper(char) {
				hasUppercase = true
			}

			if unicode.IsDigit(char) {
				hasNumber = true
			}

			if !unicode.IsLetter(char) && unicode.IsDigit(char) {
				hasSymbol = true
			}
		}

		return hasUppercase && hasNumber && hasSymbol
	})

	err := validate.Struct(data)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		var validationErrors []string

		for _, e := range errs {
			validationErrors = append(validationErrors, e.Translate(trans))
		}

		return validationErrors
	}

	return nil
}
