package controller

import (
	"errors"
	"github.com/insektionen/IN-TV/api"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"path/filepath"
)

type albumController struct{}

func NewAlbumController() api.Controller {
	return &albumController{}
}

func (c *albumController) HandleGet(r *http.Request) (api.ViewModel, error) {
	path := viper.GetString("paths.album_storage")
	if album := api.FromRoute("name", r); album != "" {
		path = filepath.Join(path, album)
	}

	dirEntries, err := os.ReadDir(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, api.NotFoundError()
	} else if err != nil {
		return nil, err
	}
	res := make([]string, 0, len(dirEntries))
	for _, d := range dirEntries {
		res = append(res, d.Name())
	}
	return api.JSONView(res), nil
}

func (c *albumController) HandlePost(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c *albumController) HandlePut(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c *albumController) HandleDelete(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}
