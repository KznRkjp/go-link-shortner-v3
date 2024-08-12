package main

import (
	"net/http"

	"github.com/KznRkjp/go-link-shortner-v3.git/cfg"
	"github.com/KznRkjp/go-link-shortner-v3.git/internal/app"
)

func main() {

	http.HandleFunc("/", app.MainPagePost)
	http.HandleFunc("/{id}", app.MainPageGet)
	err := http.ListenAndServe(cfg.Server, nil)
	if err != nil {
		panic(err)
	}

}
