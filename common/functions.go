package common

import (
	"strings"
)

func GetImageSource(imageLocation string) string {
	return strings.Join([]string{"/img/posts/", imageLocation}, "")
}
