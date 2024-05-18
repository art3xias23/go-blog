package common

import (
	"strings"
)

func GetImageSource(imageLocation string) string {
	return strings.Join([]string{"/img/posts/", imageLocation}, "")
}

func GetEndpoint(base string, endpoint string) string{
	return "/" + base + "/" + endpoint
}
