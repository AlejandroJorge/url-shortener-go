package repository

import (
	"database/sql"

	"github.com/AlejandroJorge/url-shortener-go/model"
	"github.com/AlejandroJorge/url-shortener-go/util"
)

type URLRepositoryImpl struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) URLRepository {
	return URLRepositoryImpl{db}
}

func (repository URLRepositoryImpl) Save(register model.URL) {
	insertSQL := "INSERT INTO URL (ORIGINAL_URL, PROPIETARY_ROUTE, VISITS) VALUES (?,?,?)"
	_, err := repository.db.Exec(insertSQL, register.OriginalURL, register.PropietaryRoute, register.Visits)
	util.PanicIfError(err)
}

func (repository URLRepositoryImpl) GetOriginalURL(shortenedPath string) string {
	querySQL := "SELECT ORIGINAL_URL FROM URL WHERE PROPIETARY_ROUTE = (?)"
	row := repository.db.QueryRow(querySQL, shortenedPath)

	var originalURL string
	err := row.Scan(&originalURL)
	util.PanicIfError(err)

	return originalURL
}
