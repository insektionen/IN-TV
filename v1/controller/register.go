package controller

import (
	"encoding/json"
	"github.com/insektionen/IN-TV/api"
	"github.com/insektionen/IN-TV/slideshow"
	"github.com/insektionen/IN-TV/v1/viewmodels"
	"net/http"
)

type registerController struct{}

func NewRegisterController() api.Controller {
	return &registerController{}
}

func (c *registerController) HandleGet(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c *registerController) HandlePost(r *http.Request) (api.ViewModel, error) {
	reg := &viewmodels.Register{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(reg)
	if err != nil {
		return nil, api.BadRequestError("Could not decode JSON")
	}
	slideshow.RegisterClient(reg)
	return api.JSONView(reg), err
}

func (c *registerController) HandlePut(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c *registerController) HandleDelete(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}
