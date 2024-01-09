package main

import (
	"net/http"

	"github.com/a-h/templ"
	comps "github.com/art3xias23/go-blog/components"
)

func main() {

	title := comps.Title("Tinolog")
	homecmp := comps.Home(nil)
	layout := comps.Layout(title, homecmp)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./components/assets/"))))
	http.Handle("/", templ.Handler(layout))
	// http.Handle("/latestPosts", serveLatestsPosts)

	http.ListenAndServe(":3000", nil)

}

func serveLatestsPosts(r *http.Request, w http.ResponseWriter) {

}
