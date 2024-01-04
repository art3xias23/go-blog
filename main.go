package main

import (
	"net/http"

	"github.com/a-h/templ"
	comps "github.com/art3xias23/go-blog/components"
)

func main() {

	title := comps.Title("Homepage")
	homecmp := comps.HomeContents()
	layout := comps.Layout(title, homecmp)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./components/assets/"))))
	http.Handle("/", templ.Handler(layout))

	http.ListenAndServe(":3000", nil)
}
