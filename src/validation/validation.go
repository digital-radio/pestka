//Package validation implements tools to validate input
package validation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	customerrors "github.com/digital-radio/pestka/src/custom_errors"
	"gopkg.in/go-playground/validator.v9"
)

//Validator gathers methods to perform validation/clean input.
type Validator struct {
}

//CleanJSON verifies if body is json and loads it to the required structure - if not possible it raises AppError. Parameter 'input' must be a pointer.
func (v *Validator) CleanJSON(body []byte, input interface{}) error {

	var validate *validator.Validate = validator.New()

	inputKind := reflect.ValueOf(input).Kind()
	if inputKind != reflect.Ptr {
		err := fmt.Errorf("input assertion error in: Validator -> CleanJson. Input is not a pointer - kind: %v!\n", inputKind)
		appError := customerrors.AppError{Err: err, Code: http.StatusInternalServerError, Message: "Internal server error"}
		return &appError
	}

	err := json.Unmarshal(body, input)
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
