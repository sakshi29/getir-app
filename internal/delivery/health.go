package delivery

import "net/http"

func (a *API) Health(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return "ok", nil
}
