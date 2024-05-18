package rssHelper

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/art3xias23/go-blog/common"
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

	var finishedBooks= common.Filter(books, func(b *Book) bool{
		fmt.Println(b.Bookshelves)
		return len(b.Bookshelves)== 0
	})

	return finishedBooks, nil
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
	BookID                   int     `csv:"Book Id"`
	Title                    string  `csv:"Title"`
	Author                   string  `csv:"Author"`
	AuthorLF                 string  `csv:"Author l-f"`
	AdditionalAuthors        string  `csv:"Additional Authors"`
	ISBN                     string  `csv:"ISBN"`
	ISBN13                   string  `csv:"ISBN13"`
	MyRating                 int     `csv:"My Rating"`
	AverageRating            float64 `csv:"Average Rating"`
	Publisher                string  `csv:"Publisher"`
	Binding                  string  `csv:"Binding"`
	NumberOfPages            int     `csv:"Number of Pages"`
	YearPublished            int     `csv:"Year Published"`
	OriginalPublicationYear  int     `csv:"Original Publication Year"`
	DateRead                 string  `csv:"Date Read"`
	DateAdded                string  `csv:"Date Added"`
	Bookshelves              string  `csv:"Bookshelves"`
	BookshelvesWithPositions string  `csv:"Bookshelves with Positions"`
	ExclusiveShelf           string  `csv:"Exclusive Shelf"`
	MyReview                 string  `csv:"My Review"`
	Spoiler                  string  `csv:"Spoiler"`
	PrivateNotes             string  `csv:"Private Notes"`
	ReadCount                int     `csv:"Read Count"`
	OwnedCopies              int     `csv:"Owned Copies"`
}
