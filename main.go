package main

import (
	"fmt"
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

	util.PanicIfError(util.RunMigration(db))

	URLRepository := repository.NewURLRepository(db)
	URLService := service.NewURLService(URLRepository)
	URLController := controller.NewURLController(URLService)

	router := mux.NewRouter()

	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.DebugLoggingMiddleware)

	router.HandleFunc("/urls", URLController.HandleShortenURL).Methods("POST")
	router.PathPrefix("/").HandlerFunc(URLController.HandleRedirectURL).Methods("GET")

	fmt.Println("App started, listening on PORT", config.GetPortString())

	util.PanicIfError(http.ListenAndServe(config.GetPortString(), router))
}
