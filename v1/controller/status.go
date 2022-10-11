package controller

import (
	"github.com/insektionen/IN-TV/api"
	"github.com/insektionen/IN-TV/slideshow"
	"github.com/insektionen/IN-TV/v1/viewmodels"
	"net/http"
)

type statusController struct{}

func NewStatusController() api.Controller {
	return &statusController{}
}

func (c *statusController) HandleGet(r *http.Request) (api.ViewModel, error) {
	runningShows := make([]string, 0, len(slideshow.RunningSlideshows))
	for _, s := range slideshow.RunningSlideshows {
		runningShows = append(runningShows, s.Slideshow.Name)
	}
	connectedScreens := make([]*viewmodels.Screen, 0, len(slideshow.ConnectedClients))
	for name, lastSeen := range slideshow.ConnectedClients {
		connectedScreens = append(connectedScreens, &viewmodels.Screen{
			Name:    name,
			LasSeen: lastSeen,
		})
	}
	res := &viewmodels.Status{
		RunningSlideshows: runningShows,
		ConnectedScreens:  connectedScreens,
	}
	return api.JSONView(res), nil
}

func (c *statusController) HandlePost(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c *statusController) HandlePut(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c *statusController) HandleDelete(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}
