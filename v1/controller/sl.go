package controller

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/insektionen/IN-TV/api"
	"github.com/insektionen/IN-TV/jobs"
)

type SLController struct{}

func NewSLController() api.Controller {
	return &SLController{}
}

// HandleGet implements api.Controller.
func (c *SLController) HandleGet(r *http.Request) (api.ViewModel, error) {

	err := jobs.CalcTimeTillDeparture(jobs.SLDataSimple)
	if err != nil {
		return nil, api.InternalServerError(err)
	}
	
	// Marshal the struct into JSON.
    jsonBytes, err := json.Marshal(jobs.SLDataSimple)
    if err != nil {
        return nil, api.InternalServerError(err)
    }
	
	return api.JSONView(jsonBytes), nil
}

// HandleDelete implements api.Controller.
func (*SLController) HandleDelete(r *http.Request) (api.ViewModel, error) {
	// panic("unimplemented")
	return nil, api.MethodNotAllowedError() // Returning method not allowed bc there is no need for Delete
}

// HandlePost implements api.Controller.
func (*SLController) HandlePost(r *http.Request) (api.ViewModel, error) {
	// panic("unimplemented")
	return nil, api.MethodNotAllowedError() // Returning method not allowed bc there is no need for Post
}

// HandlePut implements api.Controller.
func (*SLController) HandlePut(r *http.Request) (api.ViewModel, error) {
	// panic("unimplemented")
	return nil, api.MethodNotAllowedError() // Returning method not allowed bc there is no need for Put
}