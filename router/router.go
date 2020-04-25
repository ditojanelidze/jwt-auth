package router

import (
	"github.com/ditojanelidze/jwt-auth/router/routes"
	"github.com/gorilla/mux"
)
func Router() *mux.Router{
	return routes.MountRoutes(mux.NewRouter().StrictSlash(true))
}