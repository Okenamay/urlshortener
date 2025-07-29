package handlers

import (
	"io"
	"net/http"

	"github.com/Okenamay/urlshortener/internal/app/configs"
	"github.com/Okenamay/urlshortener/internal/app/database"
	"github.com/Okenamay/urlshortener/internal/app/models"
	"github.com/Okenamay/urlshortener/internal/app/services"
)

// Выбираем функцию для обработки в зависимости от метода:
func AutoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		WrapURL(w, r)
	case http.MethodGet:
		UnwrapURL(w, r)
	default:
		http.Error(w, ErrorMethodNowAllowed.Error(), http.StatusMethodNotAllowed)
	}
}

// Для метода POST свернём URL в короткий:
func WrapURL(w http.ResponseWriter, r *http.Request) {

	queryBody, readErr := io.ReadAll(r.Body)
	if readErr != nil {
		http.Error(w, ErrorServer.Error(), http.StatusInternalServerError)
		return
	}

	CheckedURL, checkErr := CheckURL(string(queryBody))

	if checkErr != nil {
		http.Error(w, checkErr.Error(), http.StatusUnprocessableEntity)
		return
	}

	originalURL := CheckedURL.String()

	shortID, genErr := services.MakeShortURL(CheckedURL)
	if genErr != nil {
		http.Error(w, genErr.Error(), http.StatusInternalServerError)
		return
	}

	newURL := fullURL(r, configs.ServerPort, shortID)

	urlRecord := models.URL{
		URL:     originalURL,
		ShortID: shortID,
	}

	if result := database.DB.Create(&urlRecord); result.Error != nil {
		http.Error(w, ErrorSaveFailed.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, newURL)
}

// Для метода GET развернём URL в изначальный и сделаем редирект:
func UnwrapURL(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Path
	if len(queryID) != models.IDSize+1 {
		http.Error(w, ErrorInvalidShortID.Error(), http.StatusNotFound)
		return
	}

	queryID = queryID[1:]

	var urlRecord models.URL
	if result := database.DB.Where("short_id = ?", queryID).First(&urlRecord); result.Error != nil {
		http.Error(w, ErrorNotInDB.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", urlRecord.URL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func fullURL(r *http.Request, port string, shortID string) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	newURL := scheme + "://localhost" + port + "/" + shortID

	return newURL
}
