package api

import "net/http"

type ViewModel interface {
	Render(w http.ResponseWriter)
}

type Controller interface {
	HandleGet(r *http.Request) (ViewModel, error)
	HandlePost(r *http.Request) (ViewModel, error)
	HandlePut(r *http.Request) (ViewModel, error)
	HandleDelete(r *http.Request) (ViewModel, error)
}

func ControllerHandler(controller Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var out ViewModel
		var err error

		switch r.Method {
		case http.MethodGet:
			out, err = controller.HandleGet(r)
		case http.MethodPost:
			out, err = controller.HandlePost(r)
		case http.MethodPut:
			out, err = controller.HandlePut(r)
		case http.MethodDelete:
			out, err = controller.HandleDelete(r)
		default:
			err = MethodNotAllowedError()
		}
		if err != nil {
			ErrorResponse(w, err)
			return
		}
		out.Render(w)
	})
}
