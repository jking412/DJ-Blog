package request

import "time"

type PostBaseResp struct {
	Id        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Tag       string
	Author    string
}

type PostIndexResp struct {
	PostBaseResp
	AvatarUrl string
}

type PostCreateReq struct {
	Title   string
	Tag     string
	Content string
}

type PostDetailResp struct {
	PostBaseResp
	Content string
	Likes   uint64
	Stared  bool
}

type PostUpdateReq struct {
	Id      uint64
	Title   string
	Tag     string
	Content string
}

type PostUpdateResp struct {
	Id      uint64
	Title   string
	Tag     string
	Content string
}

type PostSearchResp struct {
	Id    uint64
	Title string
}
