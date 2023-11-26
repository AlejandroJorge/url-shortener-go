package main

import (
	"net/http"

	"github.com/AlejandroJorge/url-shortener-go/config"
	"github.com/AlejandroJorge/url-shortener-go/controller"
	"github.com/AlejandroJorge/url-shortener-go/middleware"
	"github.com/AlejandroJorge/url-shortener-go/repository"
	"github.com/AlejandroJorge/url-shortener-go/service"
	"github.com/AlejandroJorge/url-shortener-go/util"
	"github.com/gorilla/mux"
)

func main() {

	db := config.NewDatabaseConnection()

	URLRepository := repository.NewURLRepository(db)
	URLService := service.NewURLService(URLRepository)
	URLController := controller.NewURLController(URLService)

	router := mux.NewRouter()
	router.Use(middleware.DebugLoggingMiddleware)

	router.HandleFunc("/urls", URLController.HandleShortenURL).Methods("POST")
	router.PathPrefix("/").HandlerFunc(URLController.HandleRedirectURL).Methods("GET")

	err := http.ListenAndServe("localhost:3000", router)
	util.PanicIfError(err)
}