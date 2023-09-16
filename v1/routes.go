package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/insektionen/IN-TV/api"
	"github.com/insektionen/IN-TV/v1/controller"
)

func RegisterRoutes(r *mux.Router) {
	slideshowController := controller.NewSlideshowController()
	r.Handle("/slideshow", api.ControllerHandler(slideshowController)).
		Methods(http.MethodPost, http.MethodGet)
	r.Handle("/slideshow/{name}", api.ControllerHandler(slideshowController)).
		Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	startController := controller.NewStartController()
	r.Handle("/slideshow/{name}/start", api.ControllerHandler(startController)).
		Methods(http.MethodPost)
	stopController := controller.NewStopController()
	r.Handle("/slideshow/{name}/stop", api.ControllerHandler(stopController)).
		Methods(http.MethodPost)

	statusController := controller.NewStatusController()
	r.Handle("/status", api.ControllerHandler(statusController)).
		Methods(http.MethodGet)

	registerController := controller.NewRegisterController()
	r.Handle("/register", api.ControllerHandler(registerController)).
		Methods(http.MethodPost)
	//TODO: Implement all API functions
	alumController := controller.NewAlbumController()
	//r.HandleFunc("/album", api.NotImplemented).Methods(http.MethodPost)
	r.Handle("/album", api.ControllerHandler(alumController)).Methods(http.MethodGet)
	r.Handle("/album/{album}", api.ControllerHandler(alumController)).Methods(http.MethodGet)
	//r.HandleFunc("/album/{album}", api.NotImplemented).Methods(http.MethodDelete)

	albumPictureController := controller.NewAlbumPictureController()
	r.Handle("/album/{album}", api.ControllerHandler(albumPictureController)).Methods(http.MethodPost)
	r.Handle("/album/{album}/{name}", api.ControllerHandler(albumPictureController)).Methods(http.MethodGet)
	//r.HandleFunc("/album/{album}/{name}", api.NotImplemented).Methods(http.MethodDelete)

	//r.HandleFunc("/presentation/file", api.NotImplemented).Methods(http.MethodPost)
	//r.HandleFunc("/presentation/url", api.NotImplemented).Methods(http.MethodPost)

	SLController := controller.NewSLController()
	r.Handle("/sl/info", api.ControllerHandler(SLController)).Methods(http.MethodGet)

	//r.HandleFunc("/spotify/playing", api.NotImplemented).Methods(http.MethodGet)
}
