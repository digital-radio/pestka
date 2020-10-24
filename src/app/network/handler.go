//Package network handles request to create_network (to connect to wifi) - it validates request and runs service.
package network

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/digital-radio/pestka/src/app/app"
	"github.com/digital-radio/pestka/src/utils"
	"gopkg.in/go-playground/validator.v9"
)

//Handler uses validate to validate request, uses service to perform business logic
type Handler struct {
	validate *validator.Validate
	service  *Service
}

//NewHandler creates network handler based on given validator and service.
func NewHandler(v *validator.Validate, s *Service) Handler {
	return Handler{validate: v, service: s}
}

//Expected input (request body) with validation rules.
type createNetworkRequest struct {
	Ssid     string `json:"ssid" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//Details contains domain needed to connect to the network.
type Details struct {
	Ssid     string
	Password string
}

//Create uses handler to validate request, maps input to domain and runs service to get the job done.
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		appError := app.AppError{Err: err, Code: http.BadRequestError, Message: "Bad request: failed to read the body"}
		return utils.HandleError(w, appError)
	}

	input := createNetworkRequest{}

	err = json.Unmarshal(body, &input)
	if err != nil {
		appError := app.AppError{Err: err, Code: http.BadRequestError, Message: "Bad request: not a json"}
		return utils.HandleError(w, appError)
	}

	if err := h.validate.Struct(input); err != nil {
		appError := app.AppError{Err: err, Code: http.BadRequestError, Message: "Bad request: invalid input - " + err.Error()}
		return utils.HandleError(w, appError)
	}

	//Map to domain entity.
	details := Details{
		Ssid:     input.Ssid,
		Password: input.Password,
	}

	h.service.Create(&details)
	return
}
