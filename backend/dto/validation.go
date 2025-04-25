package dto

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateRequest(s any) []string {
	errs := []string{}

	err := validate.Struct(s)
	if err != nil {
		var customMsg map[string]map[string]string
		if v, ok := s.(interface {
			ValidationMessages() map[string]map[string]string
		}); ok {
			customMsg = v.ValidationMessages()
		}
		for _, e := range err.(validator.ValidationErrors) {
			field := e.Field()
			tag := e.Tag()
			if customMsg != nil {
				if msg, ok := customMsg[field][tag]; ok {
					errs = append(errs, msg)
					continue
				}
			}
			errs = append(errs, field+" "+tag)
		}
	}

	if v, ok := s.(interface{ Validate() []string }); ok {
		errs = append(errs, v.Validate()...)
	}
	return errs
}
