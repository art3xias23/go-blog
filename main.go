package main

import (
	"context"
	"embed"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	comps "github.com/art3xias23/go-blog/components"
	"github.com/art3xias23/go-blog/domain"
	rssHelper "github.com/art3xias23/go-blog/rssHelper"
)

// go.embed secret.key
var content embed.FS

func main() {
	fileServerImg := http.FileServer(http.Dir("./components/assets/img/"))
	http.Handle("/img/", http.StripPrefix("/img/", fileServerImg))
	fileServer := http.FileServer(http.Dir("./components/styles/"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fileServer))
	fileServerjs := http.FileServer(http.Dir("./components/assets/scripts"))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fileServerjs))
	fmt.Println("Entered: main()")
	layout := comps.Layout(nil)

	http.Handle("/", templ.Handler(layout))
	http.HandleFunc("/blog", serveBlog)
	http.HandleFunc("/about", serveAbout)
	http.HandleFunc("/letterboxd", serveLetterBoxd)
	http.HandleFunc("/goodreads", serveGoodReads)
	http.HandleFunc("/letter-redirect", serveLetterRedirect)
	http.HandleFunc("/good-redirect", serveGoodRedirect)

	http.ListenAndServe("localhost:3000", nil)

}

func serveLetterRedirect(w http.ResponseWriter, r *http.Request) {
	redirectUrl := r.URL.Query().Get("url")
	w.Header().Set("HX-Redirect", redirectUrl)
	renderSenderContent(r, w, nil)
}

func serveGoodRedirect(w http.ResponseWriter, r *http.Request) {
	redirectUrl := r.URL.Query().Get("url")
	fmt.Println("Url is ", redirectUrl)
	w.Header().Set("HX-Redirect", redirectUrl)
	renderSenderContent(r, w, nil)
}

func serveLetterBoxd(w http.ResponseWriter, r *http.Request) {
	contents, err := rssHelper.GetLetterBoxdRssData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(contents.Title)
	var letterboxdView = comps.Letterboxd(contents.Items)
	renderSenderContent(r, w, letterboxdView)
}

func serveGoodReads(w http.ResponseWriter, r *http.Request) {
	contents, err := rssHelper.GetGoodReadsRssData()
	if err != nil {
		fmt.Println(err)
		return
	}
	var goodreadsView = comps.Goodreads(contents)
	// templ.Handler(blogView).ServeHTTP(w, r)
	renderSenderContent(r, w, goodreadsView)
}

// rednerSenserContent is a helper function which helps detemine if the request
// to the resource is coming from inside or outside the blog
func renderSenderContent(r *http.Request, w http.ResponseWriter, componentToRender templ.Component) {
	senderHeader := r.Header.Get("HX-Request")
	if senderHeader != "" {
		componentToRender.Render(context.Background(), w)
		return
	}
	layout := comps.Layout(componentToRender)
	layout.Render(context.Background(), w)
}

func serveAbout(w http.ResponseWriter, r *http.Request) {
	aboutView := comps.About()

	renderSenderContent(r, w, aboutView)
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
	renderSenderContent(r, w, blogView)
}
