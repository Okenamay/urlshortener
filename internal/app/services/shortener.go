package services

import (
	"crypto/rand"
	"fmt"
	"net/url"

	"github.com/Okenamay/urlshortener/internal/app/database"
	"github.com/Okenamay/urlshortener/internal/app/models"
	_ "gorm.io/gorm"
)

const (
	Charset     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"
	ShortLength = 8
)

func MakeShortURL(originalURL *url.URL) (string, error) {
	var shortID string
	var err error

	for {
		shortID, err = generateShortID()
		if err != nil {
			return "", err
		}

		// Проверяем, не существует ли уже такой ShortID в базе
		var count int64
		result := database.DB.Model(&models.URL{}).Where("short_id = ?", shortID).Count(&count)

		if result.Error != nil {
			return "", fmt.Errorf("error checking ShortID uniqueness: %w", result.Error)
		}

		if count == 0 {
			break
		}
	}

	return shortID, nil
}

// Вспомогательная функция для генерации случайной последовательности:
func generateShortID() (string, error) {
	shortID := make([]byte, ShortLength)

	if _, err := rand.Read(shortID); err != nil {
		return "", ErrorRNGFail
	}

	for i := range shortID {
		shortID[i] = Charset[shortID[i]%byte(len(Charset))]
	}

	return string(shortID), nil
}
