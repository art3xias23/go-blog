package interfaces

import (
	_ "github.com/art3xias23/go-blog/domain"
)

type PostsService interface {
	GetLatestsPosts() ([]Post, error)
}
