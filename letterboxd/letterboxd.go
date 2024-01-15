package letterboxd

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

func GetFile() (string, error) {
	url := "http://www.letterboxd.com/art3xias/rss"
	outputFile := "file.xml"

	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(fmt.Sprintln(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(outputFile, body, fs.FileMode(0644))
	if err != nil {
		return "", err
	}

	return outputFile, nil
}

func ReadFileContents(fileName string) (*Channel, error) {
	xmlContent, err := os.ReadFile(fileName)
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

type MyFeed struct {
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
}
