package data

import (
	"context"

	"github.com/KznRkjp/go-link-shortner-v3.git/cfg"
	"github.com/KznRkjp/go-link-shortner-v3.git/internal/models"
	"github.com/KznRkjp/go-link-shortner-v3.git/internal/urlgen"
)

// var ResDB []models.UrlRecord //Резидентная БД
var ResDB = make(map[string]models.URLRecord)

func SaveData(ctx context.Context, body []byte) (result string, err error) {
	url := urlgen.GenerateShortKey()
	ResDB[url] = models.URLRecord{
		OriginalURL: string(body),
		ShortURL:    url}
	result = "http://" + cfg.Server + "/" + url
	return result, nil
}
