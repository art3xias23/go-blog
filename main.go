package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	comps "github.com/art3xias23/go-blog/components"
	"github.com/art3xias23/go-blog/domain"
)

func main() {

	title := comps.Title("Tinolog")
	homecmp := comps.Home(nil)
	layout := comps.Layout(title, homecmp)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./components/assets/"))))
	http.Handle("/", templ.Handler(layout))
	http.HandleFunc("/latestPosts", serveLatestsPosts)

	http.ListenAndServe(":3000", nil)

}

func serveLatestsPosts(w http.ResponseWriter, r *http.Request) {
	mongoService, err := domain.NewMongoDbService("mongodb://localhost:27017")
	if err != nil {
		fmt.Println("Error creating MongoDB service:", err)
		return
	}

	latestPosts, err := mongoService.GetLatestsPosts()
	if err != nil {
		fmt.Println("Error retrieving latest posts:", err)
		return
	}

	// Use latestPosts as needed...
	fmt.Println("Latest Posts:", latestPosts)
}
