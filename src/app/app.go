//Package app implements http server api.
package app

import (
	"github.com/digital-radio/pestka/src/container"
	"github.com/digital-radio/pestka/src/utils"
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

//App allows to setup router.
type App struct {
	container container.Container
}

//New allows to create a new App struct outside of package app.
func New(container container.Container) App {
	return App{container}
}

//CreateRouter creates router and maps urls to functions.
func (a *App) CreateRouter() *mux.Router {
	var r = mux.NewRouter()
	r.HandleFunc("/networks", a.getNetworks).Methods(http.MethodGet)
	r.HandleFunc("/networks", a.createNetwork).Methods(http.MethodPost)
	r.HandleFunc("/", a.notFound)
	return r
}

func (a *App) notFound(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"message": "not found",
	}
	utils.Response(w, data, http.StatusNotFound)
}

func (a *App) getNetworks(w http.ResponseWriter, r *http.Request) {
	cells, err := a.container.Scan(a.container.InterfaceName)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	utils.Response(w, cells, http.StatusOK)
}

type networkDetails struct {
	Ssid string
	Password string
}

func (a *App) createNetwork(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()	

	body, err := ioutil.ReadAll(r.Body)	
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	
	var n networkDetails
	err = json.Unmarshal(body, &n)
	if err != nil {
        utils.HandleError(w, err)
		return
	}
 
	fmt.Println(n)
}