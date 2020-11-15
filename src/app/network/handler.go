//Package network handles request to create_network (to connect to wifi) - it validates request and runs service.
package network

import (
	"io/ioutil"
	"net/http"

	customerrors "github.com/digital-radio/pestka/src/custom_errors"
	"github.com/digital-radio/pestka/src/utils"
	"github.com/digital-radio/pestka/src/validation"
)

//Handler uses validate to validate request, uses service to perform business logic
type Handler struct {
	validate *validation.Validator
	service  *Service
}

//NewHandler creates network handler based on given validator and service.
func NewHandler(v *validation.Validator, s *Service) Handler {
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

//Get runs service to get networks and returns response.
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	cells, err := h.service.Get()

	if err != nil {
		appError := customerrors.AppError{Err: err, Code: http.StatusInternalServerError, Message: "Internal Server Error"}
		utils.HandleError(w, &appError)
		return
	}

	utils.Response(w, cells, http.StatusOK)
}

type responseMessage struct {
	Message string `json:"message"`
}

//Create uses handler to validate request, maps input to domain, runs service to create network and returns response.
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		appError := customerrors.AppError{Err: err, Code: http.StatusBadRequest, Message: "Bad request: failed to read the body"}
		utils.HandleError(w, &appError)
		return
	}

	input := createNetworkRequest{}

	err = h.validate.CleanJSON(body, &input)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	//Map to domain entity.
	details := Details{
		Ssid:     input.Ssid,
		Password: input.Password,
	}

	h.service.Create(&details)
	utils.Response(w, responseMessage{Message: "OK"}, http.StatusOK)
	return
}
