package routes

import (
	"github.com/ditojanelidze/jwt-auth/api/controllers"
	"net/http"
)

var AuthRoutes = []Route{
	Route{
		Uri:     "/login",
		Method:  http.MethodPost,
		Handler: controllers.Auth,
	},
	Route{
		Uri:     "/logout",
		Method:  http.MethodDelete,
		Handler: controllers.LogOut,
	},
}
