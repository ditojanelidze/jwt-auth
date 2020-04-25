package controllers

import (
	"encoding/json"
	"github.com/ditojanelidze/jwt-auth/api/services"
	"github.com/ditojanelidze/jwt-auth/response"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request){
	params, err := authParams(r)
	if err != nil {
		response.JSON(w,http.StatusUnprocessableEntity, err.Error())
		return
	}
	token, err := services.AuthService(params.Username)
	response.JSON(w, http.StatusOK, token)
}

func authParams(r *http.Request) (interface{}, error) {
	var data struct {
		Username string
		Password string
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	return data, err
}