package controllers

import (
	"encoding/json"
	"github.com/ditojanelidze/jwt-auth/api/services"
	"github.com/ditojanelidze/jwt-auth/response"
	"io/ioutil"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error_response := response.BadRequestError(err.Error())
		response.JSON(w, error_response.Code, error_response)
		return
	}

	data := struct{Username string; Password string}{}
	if err = json.Unmarshal(body, &data); err != nil{
		error_response := response.BadRequestError(err.Error())
		response.JSON(w, error_response.Code, error_response)
		return
	}

	tokens, err := services.Authenticate(data.Username,data.Password)
	if err != nil {
		error_response := response.UnautorizedError(err.Error())
		response.JSON(w, error_response.Code, error_response)
		return
	}

	response.JSON(w, http.StatusOK, tokens)
}

func LogOut(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error_response := response.BadRequestError(err.Error())
		response.JSON(w, error_response.Code, error_response)
		return
	}

	data := struct{Token string}{}
	if err = json.Unmarshal(body, &data); err != nil{
		error_response := response.BadRequestError(err.Error())
		response.JSON(w, error_response.Code, error_response)
		return
	}

	services.LogOut(data.Token)
	response.JSON(w, http.StatusOK, nil)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		error_response := response.BadRequestError(err.Error())
		response.JSON(w, error_response.Code, error_response)
		return
	}

	data := struct{Token string}{}
	if err = json.Unmarshal(body, &data); err != nil{
		error_response := response.BadRequestError(err.Error())
		response.JSON(w, error_response.Code, error_response)
		return
	}

	tokens, err := services.RefreshToken(data.Token)
	if err != nil {
		error_response := response.UnautorizedError(err.Error())
		response.JSON(w, error_response.Code, error_response)
		return
	}

	response.JSON(w, http.StatusOK, tokens)

}