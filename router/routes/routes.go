package routes

import (
	"github.com/ditojanelidze/jwt-auth/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Uri string
	Method string
	Handler func(w http.ResponseWriter, r *http.Request)
}

func MountRoutes(r *mux.Router) *mux.Router{
	for _, route := range LoadRoutes() {
		r.HandleFunc(route.Uri, middleware.WrapResponse(route.Handler)).Methods(route.Method)
	}
	return r
}

//Gathers all routes from routes/* files
func LoadRoutes() []Route {

	return []Route{}
}