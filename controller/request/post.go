package request

import "time"

type PostBaseResp struct {
	Id        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Tag       string
}
