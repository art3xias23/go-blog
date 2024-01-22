package gapi

import (
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
)

//go:embed secret.key
var content embed.FS

const GoogleBooksAPIURL = "https://www.googleapis.com/books/v1/volumes"
const GoogleBooksAPI = "AIzaSyBBYb3Mr7iaU2lfgHmb6BSwEnkYpo7-uJQ"

func GetThumbnail(isbn string) (string, error) {
	key, err := content.ReadFile("secret.key")
	if err != nil {
		fmt.Println("Error obtaining secret key:", err)
		return "", err
	}
	url := fmt.Sprintf("%s?q=isbn:%s&key=%s", GoogleBooksAPIURL, isbn, key)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making API request:", err)
		return "", err
	}

	defer response.Body.Close()
	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", response.Status)
		return "", err
	}

	var bookInfo BookInfo

	err = json.NewDecoder(response.Body).Decode(&bookInfo)
	if err != nil {
		fmt.Println("Error decoding JSON response:", err)
		return "", err
	}

	if len(bookInfo.VolumeInfo.ImageLinks.Thumbnail) > 0 {
		// Print the thumbnail URL
		return bookInfo.VolumeInfo.ImageLinks.Thumbnail, nil
	} else {
		return "", fmt.Errorf("no thumnail present")
	}
}

type BookInfo struct {
	VolumeInfo struct {
		Title               string   `json:"title"`
		Authors             []string `json:"authors"`
		ImageLinks          ImageLinks
		IndustryIdentifiers []struct {
			Type       string `json:"type"`
			Identifier string `json:"identifier"`
		} `json:"industryIdentifiers"`
	} `json:"volumeInfo"`
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}
