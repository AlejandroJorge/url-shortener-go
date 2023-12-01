package repository

import (
	"errors"

	"github.com/AlejandroJorge/url-shortener-go/customerror"
	"github.com/AlejandroJorge/url-shortener-go/model"
	"gorm.io/gorm"
)

type URLRepositoryImpl struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return URLRepositoryImpl{db: db}
}

func (repository URLRepositoryImpl) Save(register model.URL) error {
	err := repository.db.Create(&register).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			err = customerror.ErrShortenedRouteAlreadyExists
		}
		return err
	}

	return nil
}

func (repository URLRepositoryImpl) GetOriginalURL(shortenedPath string) (string, error) {
	var storedURL model.URL
	err := repository.db.Where("PROPIETARY_ROUTE = ?", shortenedPath).First(&storedURL).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = customerror.ErrNoShortenedRouteRegistered
		}
		return "", err
	}
	return storedURL.OriginalURL, nil
}

func (repository URLRepositoryImpl) GetPreviousStoredShortenedURL(originalURL string) (bool, string, error) {
	var storedURL model.URL
	err := repository.db.Where("ORIGINAL_URL = ?", originalURL).Take(&storedURL).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = customerror.ErrNoShortenedRouteRegistered
		}
		return false, "", nil
	}

	return true, storedURL.PropietaryRoute, nil
}

func (repository URLRepositoryImpl) UpdateRedirectionsCount(originalURL string) error {
	var storedURL model.URL
	err := repository.db.Where("ORIGINAL_URL = ?", originalURL).First(&storedURL).Error
	if err != nil {
		return err
	}

	storedURL.Visits++
	err = repository.db.Save(&storedURL).Error
	if err != nil {
		return err
	}

	return nil
}
