package letterboxd

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
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

	var rssFeed MyFeed
	err = xml.Unmarshal(xmlContent, &rssFeed)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return nil, err
	}

	fmt.Println("Title: ", rssFeed.Channel.Title)
	fmt.Println("Description", rssFeed.Channel.Description)
	return &rssFeed.Channel, nil
}

func GetGoodReadsRssData() (*Channel, error) {
	url := "https://www.goodreads.com/user/updates_rss/44259798"

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

	return &rssFeed.Channel, nil
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
