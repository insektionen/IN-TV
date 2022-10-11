package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/insektionen/IN-TV/api"
	"github.com/insektionen/IN-TV/v1/viewmodels"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
)

type slideshowController struct{}

func NewSlideshowController() api.Controller {
	return &slideshowController{}
}

func (c *slideshowController) HandleGet(r *http.Request) (api.ViewModel, error) {
	if name := api.FromRoute("name", r); name != "" {
		filePath := filepath.Join(viper.GetString("paths.slideshow_storage"), fmt.Sprintf("%s.json", name))
		// f is closed after the api.ReaderView has read it
		f, err := os.Open(filePath)
		if errors.Is(err, os.ErrNotExist) {
			return nil, api.NotFoundError()
		} else if err != nil {
			return nil, err
		}
		return api.ReaderView(f), nil
	}

	files, err := os.ReadDir(viper.GetString("paths.slideshow_storage"))
	if err != nil {
		return nil, err
	}
	res := make([]*viewmodels.Slideshow, 0, len(files))
	for _, ls := range files {
		s := &viewmodels.Slideshow{}
		filePath := filepath.Join(viper.GetString("paths.slideshow_storage"), ls.Name())
		f, _ := os.Open(filePath)
		_ = json.NewDecoder(f).Decode(s)
		res = append(res, s)
		_ = f.Close()
	}

	return api.JSONView(res), nil
}

func (c *slideshowController) HandlePost(r *http.Request) (api.ViewModel, error) {
	s := &viewmodels.Slideshow{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		return nil, api.BadRequestError("Could not decode JSON")
	}
	jsData, _ := json.Marshal(s)
	filePath := fmt.Sprintf("%s/%s.json", viper.GetString("paths.slideshow_storage"), s.Name)
	if _, err := os.Stat(filePath); err == nil {
		return nil, api.BadRequestError("A slideshow with that name already exists.")
	}
	err = os.WriteFile(filePath, jsData, 0600)
	if err != nil {
		return nil, err
	}
	return api.JSONView(s), nil
}

func (c *slideshowController) HandlePut(r *http.Request) (api.ViewModel, error) {
	s := &viewmodels.Slideshow{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		return nil, api.BadRequestError("Could not decode JSON")
	}
	jsData, _ := json.Marshal(s)
	filePath := filepath.Join(viper.GetString("paths.slideshow_storage"), fmt.Sprintf("%s.json", s.Name))
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return nil, api.NotFoundError()
	}
	err = os.WriteFile(filePath, jsData, 0600)
	if err != nil {
		return nil, err
	}
	return api.JSONView(s), nil
}

func (c *slideshowController) HandleDelete(r *http.Request) (api.ViewModel, error) {
	name := api.FromRoute("name", r)
	filePath := fmt.Sprintf("%s/%s.json", viper.GetString("paths.slideshow_storage"), name)
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return nil, api.NotFoundError()
	}
	err := os.Remove(filePath)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
