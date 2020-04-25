package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}){
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println("Eroor returning response")
	}
}