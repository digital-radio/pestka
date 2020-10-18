//Package netowrk handles request to create_network ((to connect to wifi) - it runs validation and service.
package network

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/digital-radio/pestka/src/container"
	"github.com/digital-radio/pestka/src/utils"
	"gopkg.in/go-playground/validator.v9"
)

type networkHandler struct {
	validate  *validator.Validate
	service   *Service
	container *container.Container
}

func NewNetworkHandler(v *validator.Validate, s *Service, c *container.Container) networkHandler {
	return networkHandler{validate: v, service: s, container: c}

}

//Expected input (request body) with validation rules.
type createNetworkRequest struct {
	Ssid     string `json:"ssid" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//NetworkDetails contains domain details needed to create a new network.
type NetworkDetails struct {
	Ssid     string
	Password string
}

func (h *networkHandler) Create(w http.ResponseWriter, r *http.Request) {

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
	n := NetworkDetails{
		Ssid:     input.Ssid,
		Password: input.Password,
	}

	//Finally pass the entity into service to get the job done.
	h.service.Create(&n)
	return
}
