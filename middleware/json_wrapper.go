package middleware

import (
	"log"
	"net/http"
)

func WrapResponse(handler http.HandlerFunc) http.HandlerFunc {
	return restResponse(responseLogger(handler))
}

//Turn responses into json
func restResponse(handler http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		handler(w,r)
	}
}

//Add logging functionality to every request
func responseLogger(handler http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		log.Printf("\n%s %s%s %s",r.Method, r.Host, r.RequestURI, r.Proto )
		handler(w,r)
	}
}