package main

import (
	"context"
	"embed"
	"io/fs"
	"os"

	"log"
	// "net"
	"net/http"
	"time"

	"github.com/a-h/templ"
	comps "github.com/art3xias23/go-blog/components"
	"github.com/art3xias23/go-blog/domain"
	rssHelper "github.com/art3xias23/go-blog/rssHelper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// go:embed secret.key
var content embed.FS

//go:embed components/assets/img/*
var images embed.FS

//go:embed components/assets/scripts/*
var scripts embed.FS

func main() {
	logFile, err := os.OpenFile("/home/tino/go/go-blog/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
		logFile, err = os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	log.Println("Enter App")
	fileServerIcon := http.FileServer(http.Dir("./components/assets/icon/"))
	http.Handle("/icon/", http.StripPrefix("/icon/", fileServerIcon))

	staticFs, err := fs.Sub(images, "components/assets/img")
	if err != nil {
		log.Println("Could not read static img files")
		log.Println(err)
	}
	fileServerImg := http.FileServer(http.FS(staticFs))
	http.Handle("/img/", http.StripPrefix("/img/", fileServerImg))

	fileServer := http.FileServer(http.Dir("./components/styles/"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fileServer))

	staticFsJs, err := fs.Sub(scripts, "components/assets/scripts")
	fileServerjs := http.FileServer(http.FS(staticFsJs))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", fileServerjs))

	http.HandleFunc("/", servePosts)
	http.HandleFunc("/posts", servePosts)
	http.HandleFunc("/posts/new", servePostNew)
	http.HandleFunc("/posts/{id}", servePost)
	http.HandleFunc("/tags/{tag}", serveTag)
	http.HandleFunc("/about", serveAbout)
	http.HandleFunc("/tag-add", serveTagAdd)
	http.HandleFunc("/letterboxd", serveLetterBoxd)
	http.HandleFunc("/goodreads", serveGoodReads)
	http.HandleFunc("/letter-redirect", serveLetterRedirect)
	http.HandleFunc("/good-redirect", serveGoodRedirect)

	log.Println("Loaded on :3000")
	http.ListenAndServe(":3000", nil)

	// addr := net.JoinHostPort("::", "8100")
	// server := &http.Server{Addr: addr}
	// log.Fatalln(server.ListenAndServe())

}

func serveTagAdd(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println("Error parsing form in TagAdd")
		log.Println(err)
	}
	tag := r.FormValue("tag")

	tagComponent := comps.Tag(tag)
	renderSenderContent(r, w, tagComponent)

}
func servePostNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleNewPostGet(w, r)
	case "POST":
		handleNewPostPost(w, r)
	}
}

func handleNewPostGet(w http.ResponseWriter, r *http.Request) {
	newPostView := comps.PostNew()
	renderSenderContent(r, w, newPostView)
}

func handleNewPostPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {

		log.Println("Error parsing form in handleNewPostPost")
		log.Println(err)
	}

	title := r.FormValue("title")
	desc := r.FormValue("desc")
	content := r.FormValue("content")
	imgurl := r.FormValue("imgurl")
	tagList := r.Form["tagItems"]

	tags := make([]string, 0)

	for _, tag := range tagList {

		tags = append(tags, tag)
	}

	for cc, tagg := range tagList {
		log.Printf("Tag%d: %s\n", cc, tagg)
	}

	mongoService, err := domain.NewMongoDbService()
	if err != nil {

		log.Printf("error in mongo: ", err)
	}
	post := domain.Post{
		ID:            primitive.NewObjectID(),
		Title:         title,
		Description:   desc,
		Content:       content,
		Tags:          tags,
		Author:        "Konstantin Milchev",
		ImageLocation: imgurl,
		Created:       time.Now(),
	}

	_, err = mongoService.InsertPost(post)
	if err != nil {

		log.Println("error in inserting", err)
	}

}

func serveLetterRedirect(w http.ResponseWriter, r *http.Request) {
	redirectUrl := r.URL.Query().Get("url")
	w.Header().Set("HX-Redirect", redirectUrl)
	renderSenderContent(r, w, nil)
}

func serveGoodRedirect(w http.ResponseWriter, r *http.Request) {
	redirectUrl := r.URL.Query().Get("url")
	w.Header().Set("HX-Redirect", redirectUrl)
	renderSenderContent(r, w, nil)
}

func serveLetterBoxd(w http.ResponseWriter, r *http.Request) {
	contents, err := rssHelper.GetLetterBoxdRssData()
	if err != nil {
		log.Printf("Error in serveLetterBoxd:", err)
		return
	}
	var letterboxdView = comps.Letterboxd(contents.Items)
	renderSenderContent(r, w, letterboxdView)
}

func serveGoodReads(w http.ResponseWriter, r *http.Request) {
	contents, err := rssHelper.GetGoodReadsRssData()
	if err != nil {
		log.Printf("Error in serveGoodReads:", err)
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
func servePost(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	mongoService, err := domain.NewMongoDbService()
	if err != nil {
		log.Printf("Error creating MongoDB service: ", err)
		return
	}

	post, err := mongoService.GetPostById(idString)
	if err != nil {
		log.Printf("Error obtaining post:", err)
		return
	}
	postView := comps.Post(post)

	renderSenderContent(r, w, postView)
}

func serveTag(w http.ResponseWriter, r *http.Request) {

	// mongocs := "mongodb://172.28.224.1:27017/"
	tag := r.PathValue("tag")

	mongoService, err := domain.NewMongoDbService()
	if err != nil {
		log.Printf("Error creating MongoDB service:", err)
		return
	}

	posts, err := mongoService.GetPostsByTag(tag)
	if err != nil {
		log.Printf("Error obtaining posts from tag:", err)
		return
	}
	postView := comps.Posts(posts)
	blogView := comps.PostsMain(postView)

	renderSenderContent(r, w, blogView)
}

func servePosts(w http.ResponseWriter, r *http.Request) {
	// mongocs := "mongodb://172.28.224.1:27017/
	log.Println("Started serving posts")

	mongoService, err := domain.NewMongoDbService()
	if err != nil {
		log.Printf("Error creating MongoDB service:", err)
		return
	}
	defer mongoService.Disconnect()

	latestPosts, err := mongoService.GetPosts()
	if err != nil {
		log.Println("Error retrieving latest posts:", err)
		return
	}

	var latestPostsView = comps.Posts(latestPosts)
	var blogView = comps.PostsMain(latestPostsView)
	// templ.Handler(blogView).ServeHTTP(w, r)
	renderSenderContent(r, w, blogView)
}
