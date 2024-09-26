package validate

import (
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ru_translations "github.com/go-playground/validator/v10/translations/ru"
	"log"
	"reflect"
)

var fieldTranslations = map[string]string{
	"Name":     "Имя",               // Translated field name for 'Name'
	"Email":    "Электронная почта", // Translated field name for 'Email'
	"Password": "Пароль",
	"Login":    "Логин",
	"Role":     "Должность",
}

var Validator *validator.Validate
var Trans ut.Translator

func InitValidator() {
	lang := ru.New()
	uni := ut.New(lang, lang)
	Trans, _ = uni.GetTranslator(lang.Locale())

	Validator = validator.New()

	// Register Russian translations for validator
	if err := ru_translations.RegisterDefaultTranslations(Validator, Trans); err != nil {
		log.Fatalf("Failed to register Russian translations: %v", err)
	}

	Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		return fieldTranslations[field.Name]
	})
}
