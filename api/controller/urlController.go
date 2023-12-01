package controller

import (
	"net/http"
	"strings"

	"github.com/AlejandroJorge/url-shortener-go/customerror"
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

	shortenURLRes, err := controller.urlService.ShortenURL(shortenURLReq)
	shortenURLRes.ShortenedURL = util.ConstructURL(r.Host, shortenURLRes.ShortenedURL)

	var webResponse data.WebResponse

	switch err {
	case customerror.ErrShortenedRouteAlreadyExists:
		webResponse = data.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Shortened route already exists",
			Data:   nil,
		}
	case customerror.ErrInvalidOriginalURL:
		webResponse = data.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Invalid format of the url",
			Data:   nil,
		}
	case nil:
		webResponse = data.WebResponse{
			Code:   http.StatusAccepted,
			Status: "Ok",
			Data:   shortenURLRes,
		}
	default:
		webResponse = data.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal server error",
			Data:   nil,
		}
	}

	util.WriteResponseHeaders(w, webResponse)
	util.WriteResponseBody(w, webResponse)
}

func (controller URLController) HandleRedirectURL(w http.ResponseWriter, r *http.Request) {
	shortenedPath := r.URL.Path
	shortenedPath = strings.TrimPrefix(shortenedPath, "/")

	originalURL, err := controller.urlService.GetOriginalURL(shortenedPath)

	var webResponse data.WebResponse
	switch err {
	case customerror.ErrNoShortenedRouteRegistered:
		webResponse = data.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Shortened route not found",
			Data:   nil,
		}
	case nil:
		http.Redirect(w, r, originalURL, http.StatusPermanentRedirect)
		return
	default:
		webResponse = data.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal server error",
			Data:   nil,
		}
	}

	util.WriteResponseHeaders(w, webResponse)
	util.WriteResponseBody(w, webResponse)
}
