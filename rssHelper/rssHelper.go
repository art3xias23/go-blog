package rssHelper

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
)

func GetLetterBoxdRssData() (*Channel, error) {
	url := "http://www.letterboxd.com/art3xias/rss"

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(fmt.Sprintln(resp.StatusCode))
	}

	xmlContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed MyLetterBoxdFeed
	err = xml.Unmarshal(xmlContent, &rssFeed)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return nil, err
	}

	fmt.Println("Title: ", rssFeed.Channel.Title)
	fmt.Println("Description", rssFeed.Channel.Description)
	return &rssFeed.Channel, nil
}

func GetGoodReadsRssData() ([]*Book, error) {

	in, err := os.Open("goodreads.csv")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	books := []*Book{}

	if err := gocsv.UnmarshalFile(in, &books); err != nil {
		panic(err)
	}
	for _, book := range books {
		fmt.Println("Hello, ", book.Title)
	}

	return books, nil
}

type MyLetterBoxdFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string  `xml:"title"`
	Description string  `xml:"descrietion"`
	Items       []Itemm `xml:"item"`
}

type Itemm struct {
	Title        string `xml:"filmTitle"`
	Year         string `xml:"filmYear"`
	MemberRating string `xml:"memberRating"`
	Description  string `xml:"description"`
	Url          string `xml:"link"`
}

type Book struct {
	BookID                   int     `index:"0"`
	Title                    string  `index:"1"`
	Author                   string  `index:"2"`
	AuthorLF                 string  `index:"3"`
	AdditionalAuthors        string  `index:"4"`
	ISBN                     string  `index:"5"`
	ISBN13                   string  `index:"6"`
	MyRating                 int     `index:"7"`
	AverageRating            float64 `index:"8"`
	Publisher                string  `index:"9"`
	Binding                  string  `index:"10"`
	NumberOfPages            int     `index:"11"`
	YearPublished            int     `index:"12"`
	OriginalPublicationYear  int     `index:"13"`
	DateRead                 string  `index:"14"`
	DateAdded                string  `index:"15"`
	Bookshelves              string  `index:"16"`
	BookshelvesWithPositions string  `index:"17"`
	ExclusiveShelf           string  `index:"18"`
	MyReview                 string  `index:"19"`
	Spoiler                  string  `index:"20"`
	PrivateNotes             string  `index:"21"`
	ReadCount                int     `index:"22"`
	OwnedCopies              int     `index:"23"`
}
