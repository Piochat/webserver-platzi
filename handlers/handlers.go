package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/kokgudiel2/webserver/model"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Vos Tumadre")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Tumadre en Home")
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Error resource not found", http.StatusNotFound)
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)

	var user model.User
	err := decode.Decode(&user)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		return
	}

	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
