package request

import (
	"github.com/thedevsaddam/govalidator"
)

type ArticleCreateReq struct {
	Title         string `valid:"title" json:"title,omitempty" example:"required"`
	OriginContent string `valid:"originContent" json:"originContent,omitempty" example:"required"`
	ParseContent  string `valid:"parseContent" json:"parseContent,omitempty" example:"required"`
	ImgUrl        string `json:"imgUrl,omitempty"`
	UserId        uint32 `json:"userId,omitempty"`
}

type PaginationReq struct {
	PageNum  int `json:"pageNum" example:"1"`
	PageSize int `json:"pageSize" example:"10"`
}

type ArticleUpdateReq struct {
	Id            uint32 `valid:"id" json:"id,omitempty" example:"required"`
	Title         string `valid:"title" json:"title,omitempty" example:"required"`
	OriginContent string `valid:"originContent" json:"originContent,omitempty" example:"required"`
	ParseContent  string `valid:"parseContent" json:"parseContent,omitempty" example:"required"`
	ImgUrl        string `json:"imgUrl,omitempty"`
	UserId        uint32 `json:"userId,omitempty"`
}

func ValidateArticleReq(data interface{}) map[string][]string {
	rule := govalidator.MapData{
		"title":         []string{"required"},
		"originContent": []string{"required"},
		"parseContent":  []string{"required"},
	}

	msg := govalidator.MapData{
		"title": []string{
			"required:标题为必填项",
		},
		"originContent": []string{
			"required:内容为必填项",
		},
		"parseContent": []string{
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
