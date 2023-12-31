package main

import (
	"net/http"

	"github.com/a-h/templ"
	comps "github.com/art3xias23/go-blog/components"
)

func main() {
	// cssContent, err := os.ReadFile("assets/style.css")
	// if err != nil {
	// 	//
	// }
	// cssString := string(cssContent)

	title := comps.Title("Homepage")
	homecmp := comps.Home()
	layout := comps.Layout(title, homecmp)

	http.Handle("/", templ.Handler(layout))

	http.ListenAndServe(":3000", nil)
}
