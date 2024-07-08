package handler

import (
	"api-template/handler/hello"
	"github.com/Infinitare/types-template/requests"
	"github.com/gorilla/mux"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"net/http"
)

func Add() *muxtrace.Router {

	router := muxtrace.NewRouter().StrictSlash(true)
	router.Use(mux.CORSMethodMiddleware(router.Router))
	router.Use(requests.HandleCORS)

	router.HandleFunc("/BYsDXTr0gEBWK35kQf0i", nil).Methods(http.MethodGet)
	router.HandleFunc("/hello", hello.Handle).Methods(http.MethodPost, http.MethodOptions)

	router.NotFoundHandler = requests.NotFoundHandler
	return router

}
