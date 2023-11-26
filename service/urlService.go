package service

import (
	"github.com/AlejandroJorge/url-shortener-go/config"
	"github.com/AlejandroJorge/url-shortener-go/data"
	"github.com/AlejandroJorge/url-shortener-go/logic"
	"github.com/AlejandroJorge/url-shortener-go/model"
	"github.com/AlejandroJorge/url-shortener-go/repository"
	"github.com/AlejandroJorge/url-shortener-go/util"
)

type URLService struct {
	urlRepository repository.URLRepository
}

func NewURLService(URLrepository repository.URLRepository) URLService {
	return URLService{URLrepository}
}

func (service URLService) ShortenURL(req data.ShortenURLRequest) data.ShortenURLResponse {

	originalURL := req.OriginalURL

	shortenedPath := logic.TransformURL(originalURL)

	newURLRegister := model.URL{
		OriginalURL:     originalURL,
		PropietaryRoute: shortenedPath,
		Visits:          0,
	}

	service.urlRepository.Save(newURLRegister)

	return data.ShortenURLResponse{
		ShortenedURL: util.ConstructURL(config.Domain, config.Port, shortenedPath),
	}

}

func (service URLService) GetOriginalURL(shortenedPath string) string {
	originalURL := service.urlRepository.GetOriginalURL(shortenedPath)
	return originalURL
}
