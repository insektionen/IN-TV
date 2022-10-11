package controller

import (
	"github.com/insektionen/IN-TV/api"
	"net/http"
)

type albumPictureController struct{}

func NewAlbumPictureController() api.Controller {
	return &albumPictureController{}
}

func (c *albumPictureController) HandleGet(r *http.Request) (api.ViewModel, error) {
	album := api.FromRoute("album", r)
	if album == "" {
		return nil, api.NotFoundError()
	}
	name := api.FromRoute("name", r)
	if name == "" {
		return nil, api.NotFoundError()
	}
	//picturePath := filepath.Join(viper.GetString("paths.album_storage"), album, name)
	//f, err := os.Open(picturePath)
	//if errors.Is(err, os.ErrNotExist) {
	//	return nil, api.NotFoundError()
	//} else if err != nil {
	//	return nil, err
	//}
	// TODO: Create c buffered reader ViewModel
	//_, _ = io.Copy(w, f)
	return nil, nil
}

func (c *albumPictureController) HandlePost(r *http.Request) (api.ViewModel, error) {
	album := api.FromRoute("album", r)
	if album == "" {
		return nil, api.NotFoundError()
	}
	//picturePath := filepath.Join(viper.GetString("paths.album_storage"), album, name)
	//TODO: Save all uploaded pictures (Form multipart)
	return nil, nil
}

func (c *albumPictureController) HandlePut(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}

func (c *albumPictureController) HandleDelete(r *http.Request) (api.ViewModel, error) {
	//TODO implement me
	panic("implement me")
}
