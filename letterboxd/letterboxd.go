package letterboxd

import (
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
