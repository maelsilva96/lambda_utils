package lambda_utils

import (
	"fmt"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptBrTranslations "github.com/go-playground/validator/v10/translations/pt_BR"
	"github.com/maelsilva96/lambda_utils/entities"
	"log"
	"reflect"
	"strings"
)

var validate *validator.Validate
var trans ut.Translator

func init() {
	ptBr := pt_BR.New()
	uni := ut.New(ptBr, ptBr)
	trans, _ = uni.GetTranslator("pt_BR")
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		jsonName := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if jsonName == "-" {
			return ""
		}
		name := strings.SplitN(fld.Tag.Get("name"), ",", 2)[0]
		return fmt.Sprintf("%s;%s", jsonName, name)
	})
	err := ptBrTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatalln(err)
	}
}

func Valid(entity interface{}) []entities.FieldError {
	var fieldErrors []entities.FieldError
	err := validate.Struct(entity)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, m := range errs {
			fieldName := m.Field()
			dataName := strings.Split(fieldName, ";")
			jsonName := dataName[0]
			name := jsonName
			if len(dataName) > 1 {
				name = dataName[1]
			}
			fieldErrors = append(fieldErrors, entities.FieldError{
				Field:        dataName[0],
				ErrorMessage: strings.Replace(m.Translate(trans), m.Field(), name, -1),
			})
		}
		return fieldErrors
	}
	return fieldErrors
}
