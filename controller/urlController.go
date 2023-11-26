package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/AlejandroJorge/url-shortener-go/data"
	"github.com/AlejandroJorge/url-shortener-go/service"
	"github.com/AlejandroJorge/url-shortener-go/util"
)

type URLController struct {
	urlService service.URLService
}

func NewURLController(URLService service.URLService) *URLController {
	return &URLController{URLService}
}

func (controller URLController) HandleShortenURL(w http.ResponseWriter, r *http.Request) {
	shortenURLReq := data.ShortenURLRequest{}
	util.ReadRequestBody(r, &shortenURLReq)

	shortenURLRes := controller.urlService.ShortenURL(shortenURLReq)

	webResponse := data.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   shortenURLRes,
	}

	util.WriteResponseBody(w, webResponse)
}

func (controller URLController) HandleRedirectURL(w http.ResponseWriter, r *http.Request) {
	shortenedPath := r.URL.Path
	shortenedPath = strings.TrimPrefix(shortenedPath, "/")
	fmt.Println("Path:", shortenedPath)
	originalURL := controller.urlService.GetOriginalURL(shortenedPath)
	fmt.Println("OriginalURL:", originalURL)

	http.Redirect(w, r, originalURL, http.StatusPermanentRedirect)
}
