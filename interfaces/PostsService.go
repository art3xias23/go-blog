package interfaces

type PostsService interface {
	GetLatestsPosts() ([]Posts, error)
}
