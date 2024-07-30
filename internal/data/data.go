package data

import (
	"context"

	"github.com/KznRkjp/go-link-shortner-v3.git/internal/models"
	"github.com/KznRkjp/go-link-shortner-v3.git/internal/urlgen"
)

// var ResDB []models.UrlRecord //Резидентная БД
var ResDB = make(map[string]models.UrlRecord)

func SaveData(ctx context.Context, body []byte) (result string, err error) {
	url := urlgen.GenerateShortKey()
	ResDB[url] = models.UrlRecord{string(body), url}
	result = "http://localhost:8080/" + url
	return result, nil
}
