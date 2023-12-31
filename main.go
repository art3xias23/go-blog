package main

import (
	"net/http"

	"github.com/a-h/templ"

	"github.com/art3xias23/go-blog/templates"
)

func main() {
	title := templates.Title("Homepage")
	homecmp := templates.Home()
	layout := templates.Layout(title, homecmp)

	http.Handle("/", templ.Handler(layout))

	http.ListenAndServe(":3000", nil)
}
