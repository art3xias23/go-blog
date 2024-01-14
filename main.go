package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	comps "github.com/art3xias23/go-blog/components"
	"github.com/art3xias23/go-blog/domain"
)

func main() {

	fmt.Println("Entered: main()")
	title := comps.Title("Tinolog")
	layout := comps.Layout(title, nil)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./components/assets/"))))
	http.Handle("/", templ.Handler(layout))
	// http.HandleFunc("/blog", serveBlog)
	// http.Handle("/blog", templ.Handler(serveBlogComponent()))
	http.HandleFunc("/blog", serveBlog)
	http.HandleFunc("/about", serveAbout)

	http.ListenAndServe(":3000", nil)

}

func serveAbout(w http.ResponseWriter, r *http.Request) {
	aboutView := comps.About()

	aboutView.Render(context.Background(), w)
}

func serveBlog(w http.ResponseWriter, r *http.Request) {
	mongocs := "mongodb://172.28.224.1:27017/"

	mongoService, err := domain.NewMongoDbService(mongocs)
	if err != nil {
		fmt.Println("Error creating MongoDB service:", err)
		return
	}
	defer mongoService.Disconnect()

	latestPosts, err := mongoService.GetPosts()
	if err != nil {
		fmt.Println("Error retrieving latest posts:", err)
		return
	}

	var latestPostsView = comps.LatestPosts(latestPosts)
	var blogView = comps.Blog(latestPostsView)
	// templ.Handler(blogView).ServeHTTP(w, r)
	blogView.Render(context.Background(), w)
	if err != nil {
		fmt.Println("Error retrieving latest posts:", err)
		return
	}

	// Use latestPosts as needed...
	fmt.Println("Latest Posts:", latestPosts)
}
