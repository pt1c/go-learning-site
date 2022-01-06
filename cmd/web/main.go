package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pt1c/go-learning-site/pkg/config"
	"github.com/pt1c/go-learning-site/pkg/handlers"
	"github.com/pt1c/go-learning-site/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	app.TemplateCache, _ = render.CreateTemplateCache()

	http.HandleFunc("/", handlers.IndexPage)
	http.HandleFunc("/about", handlers.AboutPage)

	http.HandleFunc("/test", func(rw http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(rw, "Hello fuckers")
		if err != nil {
			log.Println(err)

		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	_ = http.ListenAndServe(portNumber, nil)
}
