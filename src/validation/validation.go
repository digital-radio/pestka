//Package validation implements tools to validate input
package validation

import (
	"encoding/json"
	"errors"

	"gopkg.in/go-playground/validator.v9"
)

//Validator gathers methods to perform validation/clean input.
type Validator struct {
}

//ParseAndValidateJSON verifies if body is json and loads it to the required structure - if not possible it raises AppError. Parameter 'input' must be a pointer.
func (v *Validator) ParseAndValidateJSON(body []byte, input interface{}) error {

	var validate *validator.Validate = validator.New()

	err := json.Unmarshal(body, input)
	if err != nil {
		return errors.New("not a json")
	}

	if err := validate.Struct(input); err != nil {
		return errors.New("invalid input - " + err.Error())
	}
	return nil
}
