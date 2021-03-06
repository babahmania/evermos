package middlewares

import (
	"net/http"

	"evermos/dto"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(schema interface{}) []*dto.ApiError {
	var errors []*dto.ApiError
	validate := validator.New()
	err := validate.Struct(schema)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element dto.ApiError
			element.Status = http.StatusUnprocessableEntity
			element.Code = err.StructNamespace()
			element.Title = err.Tag()
			element.Detail = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
