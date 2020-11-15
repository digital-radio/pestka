

//Package app implements http server api.
package app

import (
	"net/http"

	"github.com/digital-radio/pestka/src/app/network"
	"github.com/digital-radio/pestka/src/container"
	"github.com/digital-radio/pestka/src/utils"
	"github.com/digital-radio/pestka/src/validation"
	"github.com/gorilla/mux"
)

//App allows to setup router.
type App struct {
	container container.Container
}

//New allows to create a new App struct outside of package app.
func New(container container.Container) App {
	return App{container}
}

//CreateRouter creates router and maps urls to handlers.
func (a *App) CreateRouter() *mux.Router {
	var r = mux.NewRouter()

	v := validation.Validator{}
	s := network.NewService(&a.container)
	nh := network.NewHandler(&v, &s)

	r.HandleFunc("/networks", nh.Get).Methods(http.MethodGet)
	r.HandleFunc("/networks", nh.Create).Methods(http.MethodPost)
	r.HandleFunc("/", a.notFound)
	return r
}

func (a *App) notFound(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"message": "not found",
	}
	utils.Response(w, data, http.StatusNotFound)
}
