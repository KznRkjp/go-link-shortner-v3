package app

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/KznRkjp/go-link-shortner-v3.git/internal/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainPagePost(t *testing.T) {
	type want struct {
		code              int
		responsePrefix    string
		responseSuffixLen int
		contentType       string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code:              201,
				responsePrefix:    "http://",
				responseSuffixLen: 8,
				contentType:       "text/plain",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", nil)
			// создаем новый recorder
			w := httptest.NewRecorder()
			MainPagePost(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, tt.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			// assert.Equal(t, tt.want.response, string(resBody))
			assert.Equal(t, true, strings.HasPrefix(string(resBody), tt.want.responsePrefix))
			assert.Equal(t, tt.want.responseSuffixLen, utf8.RuneCountInString(strings.Split(string(resBody), "/")[len(strings.Split(string(resBody), "/"))-1]))
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}

func TestMainPageGet(t *testing.T) {
	type want struct {
		code         int
		responseText string
		contentType  string
	}
	tests := []struct {
		name string
		want want
	}{
		{
			name: "positive test #1",
			want: want{
				code:         307,
				responseText: "http://mail.ru",
				contentType:  "text/plain",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, _ := data.SaveData(context.Background(), []byte(tt.want.responseText))
			fmt.Println(path)
			fmt.Println(data.ResDB)
			request := httptest.NewRequest(http.MethodGet, path, nil)
			// создаем новый recorder
			w := httptest.NewRecorder()
			MainPageGet(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, tt.want.code, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)
			fmt.Println(string(resBody))
			assert.Equal(t, tt.want.responseText, string(resBody))
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
