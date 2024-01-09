package types

import "time"

type Post struct {
	Author       string
	PostLocation string
	Created      time.Time
}
