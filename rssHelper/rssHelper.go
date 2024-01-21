package rssHelper

import (
	csv "encoding/csv"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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

func GetGoodReadsRssData() (*[]Book, error) {
	url := "https://www.goodreads.com/review_porter/export/44259798/goodreads_export.csv"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	req.Header.Set("Accept", "text/csv")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(fmt.Sprintln(resp.StatusCode))
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/csv" {
		fmt.Println("Error: Unexpected content type:", contentType)
		return nil, err
	}

	r := csv.NewReader(resp.Body)
	var books []Book

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}
	cnt, err := io.ReadAll(resp.Body)
	fmt.Println(cnt)
	for {
		item, err := r.Read()
		log.Println(item)
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		bookID, _ := strconv.Atoi(record[0])
		myRating, _ := strconv.Atoi(record[7])
		averageRating, _ := strconv.ParseFloat(record[8], 64)
		numberOfPages, _ := strconv.Atoi(record[11])
		yearPublished, _ := strconv.Atoi(record[12])
		originalPublicationYear, _ := strconv.Atoi(record[13])
		readCount, _ := strconv.Atoi(record[22])
		ownedCopies, _ := strconv.Atoi(record[23])

		book := Book{
			BookID:                   bookID,
			Title:                    record[1],
			Author:                   record[2],
			AuthorLF:                 record[3],
			AdditionalAuthors:        record[4],
			ISBN:                     record[5],
			ISBN13:                   record[6],
			MyRating:                 myRating,
			AverageRating:            averageRating,
			Publisher:                record[9],
			Binding:                  record[10],
			NumberOfPages:            numberOfPages,
			YearPublished:            yearPublished,
			OriginalPublicationYear:  originalPublicationYear,
			DateRead:                 record[14],
			DateAdded:                record[15],
			Bookshelves:              record[16],
			BookshelvesWithPositions: record[17],
			ExclusiveShelf:           record[18],
			MyReview:                 record[19],
			Spoiler:                  record[20],
			PrivateNotes:             record[21],
			ReadCount:                readCount,
			OwnedCopies:              ownedCopies,
		}

		// Append Book to the slice
		books = append(books, book)

	}

	// Print the slice of Books
	fmt.Println(books)

	return &books, nil
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
