//Package validation implements tools to validate input
package validation

import (
	"encoding/json"
	"net/http"
	"reflect"

	customerrors "github.com/digital-radio/pestka/src/custom_errors"
	"gopkg.in/go-playground/validator.v9"
)

//Validator gathers methods to perform validation/clean input.
type Validator struct {
}

//CleanJSON verifies if body is json and loads it to the required structure - if not possible it raises AppError
func (v *Validator) CleanJSON(body []byte, input interface{}) error {

	var validate *validator.Validate = validator.New()

	//FIXME Dynamic input not working
	aType := reflect.TypeOf(input)

	elType := aType.Elem()

	inputPointer := reflect.New(elType).Interface()

	//FIXME Dynamic input not working

	err := json.Unmarshal(body, inputPointer)
	if err != nil {
		appError := customerrors.AppError{Err: err, Code: http.StatusBadRequest, Message: "Bad request: not a json"}
		return &appError
	}

	if err := validate.Struct(input); err != nil {
		appError := customerrors.AppError{Err: err, Code: http.StatusBadRequest, Message: "Bad request: invalid input - " + err.Error()}
		return &appError
	}
	return nil
}
