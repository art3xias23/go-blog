package gapi

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
	isbn = strings.TrimPrefix(isbn, "=\"")
	isbn = strings.TrimSuffix(isbn, "\"")
	if isbn == "" {
		fmt.Println("ISBN is empty")
		return "", nil
	}
	url := fmt.Sprintf("%s?q=isbn:%s&key=%s", GoogleBooksAPIURL, isbn, key)
	fmt.Println(url)
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

	if len(bookInfo.Items) == 0 {
		fmt.Printf("No items found for bookinfo")
		return "", nil
	}

	thumbNailUrl := bookInfo.Items[0].VolumeInfo.ImageLinks.Thumbnail
	if len(thumbNailUrl) > 0 {
		// Print the thumbnail URL

		imageResponse, err := http.Get(thumbNailUrl)
		if err != nil {
			fmt.Println("Error in thumbnail response")
			return "", nil
		}

		defer imageResponse.Body.Close()
		if response.StatusCode != http.StatusOK {
			fmt.Println("Error in thumbnail status code")
			return "", nil
		}

		imageData, err := io.ReadAll(imageResponse.Body)
		if err != nil {
			fmt.Println("Error reading body of image response")
			return "", nil
		}

		fileName := isbn
		err = os.WriteFile(fmt.Sprintf("%s.jpg", fileName), imageData, 0644)

		if err != nil {
			fmt.Println("Error writing file")
			return "", nil
		}
		fmt.Println("File saved successfully")
	} else {
		fmt.Println("No thumbnail present")
		return "", fmt.Errorf("no thumnail present")
	}
	return "", nil
}

type BookInfo struct {
	Kind       string `json:"kind"`
	TotalItems int    `json:"totalItems"`
	Items      []struct {
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
}

type ImageLinks struct {
	SmallThumbnail string `json:"smallThumbnail"`
	Thumbnail      string `json:"thumbnail"`
}
