package controller

import (
	"github.com/insektionen/IN-TV/api"
	"github.com/insektionen/IN-TV/slideshow"
	"net/http"
)

type stopController struct{}

func NewStopController() api.Controller {
	return &stopController{}
}

func (c stopController) HandleGet(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c stopController) HandlePost(r *http.Request) (api.ViewModel, error) {
	slideshowName := api.FromRoute("name", r)

	if sh, ok := slideshow.RunningSlideshows[slideshowName]; ok {
		sh.Stop()
	}

	return api.JSONView(map[string]string{"message": "Success"}), nil
}

func (c stopController) HandlePut(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c stopController) HandleDelete(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}
