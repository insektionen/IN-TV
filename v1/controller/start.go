package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/insektionen/IN-TV/api"
	"github.com/insektionen/IN-TV/mqtt"
	"github.com/insektionen/IN-TV/slideshow"
	"github.com/insektionen/IN-TV/v1/viewmodels"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
)

type startController struct{}

func NewStartController() api.Controller {
	return &startController{}
}

func (c startController) HandleGet(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c startController) HandlePost(r *http.Request) (api.ViewModel, error) {
	slideshowName := api.FromRoute("name", r)
	screens := make([]string, 0, 5)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&screens)
	if err != nil {
		return nil, api.BadRequestError("Could not decode JSON")
	}
	filePath := filepath.Join(viper.GetString("paths.slideshow_storage"), fmt.Sprintf("%s.json", slideshowName))
	f, err := os.Open(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return nil, api.NotFoundError()
	} else if err != nil {
		return nil, err
	}
	defer f.Close()
	sl := &viewmodels.Slideshow{}
	_ = json.NewDecoder(f).Decode(sl)
	sh := slideshow.NewShow(sl)

	for _, sName := range screens {
		topic := fmt.Sprintf("kistan/in_tv/screen/%s/slideshow", sName)
		data, _ := json.Marshal(map[string]string{"running": slideshowName})
		mqtt.Client.Publish(topic, 0, true, data)
	}
	sh.Run()

	return api.JSONView(map[string]string{"message": "Success"}), nil
}

func (c startController) HandlePut(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c startController) HandleDelete(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}
