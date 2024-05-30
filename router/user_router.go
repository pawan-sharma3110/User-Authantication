package router

import (
	"user/handler"

	"github.com/gorilla/mux"
)

func UserRouts() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/create", handler.CreateUser).Methods("POST")
	return r
}
