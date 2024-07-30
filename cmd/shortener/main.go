package main

import (
	"net/http"

	"github.com/KznRkjp/go-link-shortner-v3.git/internal/app"
)

func main() {

	http.HandleFunc("/", app.MainPagePost)
	http.HandleFunc("/{id}", app.MainPageGet)
	err := http.ListenAndServe(`:8080`, nil)
	if err != nil {
		panic(err)
	}

}
