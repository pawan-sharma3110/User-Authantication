package handler

import (
	"encoding/json"
	"net/http"
	"user/db/controls"
	modle "user/model"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("content-type", "aplication/json")
	var newUser modle.User
	json.NewDecoder(r.Body).Decode(&newUser)
	controls.InsertUser(newUser)
	json.NewEncoder(w).Encode(newUser)
}
