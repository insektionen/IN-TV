package api

import (
	"io"
	"net/http"
)

// jsonModel wrapper for rendering data as JSON
type jsonModel struct {
	data interface{}
}

// Render implements the ViewModel interface
func (m *jsonModel) Render(w http.ResponseWriter) {
	JSONResponse(w, m.data)
}

// JSONView converts obj to a json readable view model.
func JSONView(obj interface{}) ViewModel {
	return &jsonModel{
		data: obj,
	}
}

type readerViewModel struct {
	r io.ReadCloser
}

func (m *readerViewModel) Render(w http.ResponseWriter) {
	_, _ = io.Copy(w, m.r)
	_ = m.r.Close()
}

// ReaderView copies from the reader r and writes to w on render
func ReaderView(r io.ReadCloser) ViewModel {
	return &readerViewModel{
		r: r,
	}
}
