package request

import (
	"github.com/thedevsaddam/govalidator"
	"time"
)

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
	Title   string `valid:"title"`
	Tag     string
	Content string `valid:"content"`
}

type PostDetailResp struct {
	PostBaseResp
	AvatarUrl string
	Content   string
	Likes     uint64
	Stared    bool
}

type PostUpdateReq struct {
	Id      uint64
	Title   string `valid:"title"`
	Tag     string
	Content string `valid:"content"`
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

func ValidatePost(data interface{}) map[string][]string {
	rule := govalidator.MapData{
		"title":   []string{"required"},
		"content": []string{"required"},
	}

	msg := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
		},
		"content": []string{
			"required:内容为必填项",
		},
	}

	opts := govalidator.Options{
		Data:          data,
		Rules:         rule,
		TagIdentifier: "valid",
		Messages:      msg,
	}

	return govalidator.New(opts).ValidateStruct()
}
