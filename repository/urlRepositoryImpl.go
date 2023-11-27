package repository

import (
	"database/sql"

	"github.com/AlejandroJorge/url-shortener-go/customerror"
	"github.com/AlejandroJorge/url-shortener-go/model"
	"github.com/mattn/go-sqlite3"
)

type URLRepositoryImpl struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) URLRepository {
	return URLRepositoryImpl{db}
}

func (repository URLRepositoryImpl) Save(register model.URL) error {
	insertSQL := "INSERT INTO URL (ORIGINAL_URL, PROPIETARY_ROUTE, VISITS) VALUES (?,?,?)"
	_, err := repository.db.Exec(insertSQL, register.OriginalURL, register.PropietaryRoute, register.Visits)
	if err != nil {
		sqliteErr, ok := err.(sqlite3.Error)
		if ok && sqliteErr.Code == sqlite3.ErrConstraint {
			err = customerror.ErrShortenedRouteAlreadyExists
		}
		return err
	}

	return nil
}

func (repository URLRepositoryImpl) GetOriginalURL(shortenedPath string) (string, error) {
	querySQL := "SELECT ORIGINAL_URL FROM URL WHERE PROPIETARY_ROUTE = (?)"
	row := repository.db.QueryRow(querySQL, shortenedPath)

	var originalURL string
	err := row.Scan(&originalURL)
	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.ErrNoShortenedRouteRegistered
		}

		return originalURL, err
	}

	return originalURL, nil
}

func (repository URLRepositoryImpl) GetPreviousStoredShortenedURL(originalURL string) (bool, string, error) {
	querySQL := "SELECT PROPIETARY_ROUTE FROM URL WHERE ORIGINAL_URL = ?"
	row := repository.db.QueryRow(querySQL, originalURL)

	var existingShortenedRoute string
	err := row.Scan(&existingShortenedRoute)
	switch err {
	case sql.ErrNoRows:
		return false, existingShortenedRoute, nil
	case nil:
	default:
		return false, existingShortenedRoute, err
	}

	return true, existingShortenedRoute, nil
}

func (repository URLRepositoryImpl) UpdateRedirectionsCount(originalURL string) error {
	updateSQL := "UPDATE URL SET VISITS = VISITS + 1 WHERE ORIGINAL_URL = ?"
	_, err := repository.db.Exec(updateSQL, originalURL)
	if err != nil {
		return err
	}
	return nil
}
