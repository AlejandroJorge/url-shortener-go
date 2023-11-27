package service

import (
	"fmt"

	"github.com/AlejandroJorge/url-shortener-go/customerror"
	"github.com/AlejandroJorge/url-shortener-go/data"
	"github.com/AlejandroJorge/url-shortener-go/logic"
	"github.com/AlejandroJorge/url-shortener-go/model"
	"github.com/AlejandroJorge/url-shortener-go/repository"
)

type URLService struct {
	urlRepository repository.URLRepository
}

func NewURLService(URLrepository repository.URLRepository) URLService {
	return URLService{URLrepository}
}

func (service URLService) ShortenURL(req data.ShortenURLRequest) (data.ShortenURLResponse, error) {

	originalURL := req.OriginalURL

	if !logic.IsValidURL(originalURL) {
		return data.ShortenURLResponse{}, customerror.ErrInvalidOriginalURL
	}

	shortenedExists, shortenedPath, err := service.urlRepository.GetPreviousStoredShortenedURL(originalURL)
	if err != nil {
		return data.ShortenURLResponse{}, err
	}

	if shortenedExists {
		return data.ShortenURLResponse{
			ShortenedURL: shortenedPath,
		}, nil
	}

	shortenedPath = logic.TransformURL(originalURL)

	newURLRegister := model.URL{
		OriginalURL:     originalURL,
		PropietaryRoute: shortenedPath,
		Visits:          0,
	}

	err = service.urlRepository.Save(newURLRegister)
	if err == customerror.ErrShortenedRouteAlreadyExists {
		return service.ShortenURL(data.ShortenURLRequest{OriginalURL: shortenedPath})
	}
	if err != nil {
		return data.ShortenURLResponse{}, err
	}

	return data.ShortenURLResponse{
		ShortenedURL: shortenedPath,
	}, nil

}

func (service URLService) GetOriginalURL(shortenedPath string) (string, error) {
	originalURL, err := service.urlRepository.GetOriginalURL(shortenedPath)
	if err != nil {
		return originalURL, err
	}

	err = service.urlRepository.UpdateRedirectionsCount(originalURL)
	if err != nil {
		fmt.Println("Error while updating redirection count")
	}

	return originalURL, nil
}
