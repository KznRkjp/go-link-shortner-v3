package app

import (
	"io"
	"net/http"
	"strings"

	"github.com/KznRkjp/go-link-shortner-v3.git/internal/data"
)

// MainPagePost - эндпоинт с методом POST и путём /. Сервер принимает в теле запроса строку URL
// как text/plain и возвращает ответ с кодом 201 и сокращённым URL как text/plain.
func MainPagePost(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		// разрешаем только POST-запросы
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(req.Body) // достаем данные из body
	if err != nil {                   // валидация
		http.Error(res, "can't read body", http.StatusBadRequest)
		return
	}
	resultURL, err := data.SaveData(req.Context(), body)
	if err != nil {
		http.Error(res, "Error saving data", http.StatusInternalServerError)
		return
	}
	res.Header().Set("content-type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(resultURL))

}

// Эндпоинт с методом GET и путём /{id}, где id — идентификатор сокращённого URL (например, /EwHXdJfB).
// В случае успешной обработки запроса сервер возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
func MainPageGet(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		// разрешаем только GET-запросы
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	shortURL := strings.Trim(req.RequestURI, "/")
	resURLStruct, ok := data.ResDB[shortURL]
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	resURL := resURLStruct.OriginalURL
	res.Header().Set("Location", resURL)
	res.WriteHeader(http.StatusTemporaryRedirect)

}
