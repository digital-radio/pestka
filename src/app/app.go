//Package app implements http server api.
package app

import (
	"github.com/digital-radio/pestka/src/container"
	"github.com/digital-radio/pestka/src/utils"
	"github.com/gorilla/mux"
	"net/http"
)

//App allows to setup router.
type App struct {
	container container.Container
}

//New allows to creare a new App struct outside of package app.
func New(container container.Container) App {
	return App{container}
}

//CreateRouter creates router and maps urls to functions.
func (a *App) CreateRouter() *mux.Router {
	var r = mux.NewRouter()
	r.HandleFunc("/networks", a.getNetworks).Methods(http.MethodGet)
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
