package handlers

import (
	"net/url"
)

// Проверим полученный в запросе URL на корректность:
func CheckURL(reqURL string) (*url.URL, error) {
	checkedURL, err := url.ParseRequestURI(reqURL)
	if err != nil {
		return nil, ErrorInvalidURL
	}

	if checkedURL.Scheme != "https" && checkedURL.Scheme != "http" {
		return nil, ErrorHTTPS
	}

	if checkedURL.Host == "" {
		return nil, ErrorNoHost
	}

	return checkedURL, nil
}
