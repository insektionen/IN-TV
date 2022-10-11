package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

//go:embed frontend/build/*
var frontend embed.FS

type frontendHandler struct {
	staticPath string
	indexPath  string
}

func (h *frontendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	path = filepath.Join(h.staticPath, path)
	f, err := frontend.Open(path)
	if os.IsNotExist(err) {
		// File is not found, serve index
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer f.Close()
	sub, _ := fs.Sub(frontend, h.staticPath)
	http.FileServer(http.FS(sub)).ServeHTTP(w, r)
}
