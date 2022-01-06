package handlers

import (
	"net/http"

	"github.com/pt1c/go-learning-site/pkg/render"
)

func IndexPage(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "index.page.html")
}

func AboutPage(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "about.page.html")
}
