//Package network handles request to create_network (to connect to wifi) - it validates request and runs service.
package network

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/digital-radio/pestka/src/container"
	"github.com/digital-radio/pestka/src/utils"
	"gopkg.in/go-playground/validator.v9"
)

//Handler uses validate to validate request, uses service to perform business logic, uses container for other dependencies.
type Handler struct {
	validate  *validator.Validate
	service   *Service
	container *container.Container
}

//NewHandler creates network handler based on given validator, service and container.
func NewHandler(v *validator.Validate, s *Service, c *container.Container) Handler {
	return Handler{validate: v, service: s, container: c}

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
		utils.HandleError(w, err)
		return
	}

	input := createNetworkRequest{}

	err = json.Unmarshal(body, &input)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	if err := h.validate.Struct(input); err != nil {
		utils.HandleError(w, err)
		return
	}

	//Map to domain entity.
	details := Details{
		Ssid:     input.Ssid,
		Password: input.Password,
	}

	h.service.Create(&details)
	return
}
