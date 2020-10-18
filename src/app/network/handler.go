//Package handler handles request to create_network ((to connect to wifi) - it runs validation and service.
package handler

import (
	"github.com/digital-radio/pestka/src/container"
	"github.com/digital-radio/pestka/src/utils"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type networkHandler struct {
	validate *validator.Validate
	service Service
	container *container.Container
}

func NewNetworkHandler(v *validator.Validate, s Service, c *container.Container) networkHandler {
    return networkHandler{validate: v, service: s, contaier: c}
}

//Expected input (request body) with validation rules.
type createNetworkRequest struct {
	ssid string `json:"ssid" validate:"required"`
	password string `json:"password" validate:"required"`
}

//NetworkDetails contains domain details needed to create a new network.
type NetworkDetails struct {
	Ssid string
	Password string
}

func (h *networkHandler) Create(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()	

	body, err := ioutil.ReadAll(r.Body)	
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	err = json.Unmarshal(body, &input)
	if err != nil {
        utils.HandleError(w, err)
		return
	

	input := createNetworkRequest{}
	if err := h.validate.Struct(input); err != nil {
        utils.HandleError(w, err)
        return 
    }
	
	//Map to domain entity.
	n := networkDetails {
		Ssid: input.ssid,
		Password: input.password,
	}
 
	//Finally pass the entity into service to get the job done.
	return h.service.Create(n)
}
