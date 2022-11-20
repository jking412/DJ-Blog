package request

import (
	"github.com/thedevsaddam/govalidator"
)

// TODO: 处理文章请求中的分类和标签，要检查是否存在

type ArticleCreateReq struct {
	Title         string   `valid:"title" json:"title,omitempty" example:"required"`
	OriginContent string   `valid:"originContent" json:"originContent,omitempty" example:"required"`
	ParseContent  string   `valid:"parseContent" json:"parseContent,omitempty" example:"required"`
	ImgUrl        string   `json:"imgUrl,omitempty"`
	UserId        uint32   `json:"-"`
	Tags          []uint32 `json:"tags,omitempty"`
	Categories    []uint32 `json:"categories,omitempty"`
}

type PaginationReq struct {
	PageNum  int `json:"pageNum" example:"1"`
	PageSize int `json:"pageSize" example:"10"`
}

type ArticleUpdateReq struct {
	Id            uint32   `valid:"id" json:"id,omitempty" example:"required"`
	Title         string   `valid:"title" json:"title,omitempty" example:"required"`
	OriginContent string   `valid:"originContent" json:"originContent,omitempty" example:"required"`
	ParseContent  string   `valid:"parseContent" json:"parseContent,omitempty" example:"required"`
	ImgUrl        string   `json:"imgUrl,omitempty"`
	UserId        uint32   `json:"-"`
	Tags          []uint32 `json:"tags,omitempty"`
	Categories    []uint32 `json:"categories,omitempty"`
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

func ValidateUpdateArticleReq(data interface{}) map[string][]string {
	rule := govalidator.MapData{
		"id":            []string{"required", "exist:articleId"},
		"title":         []string{"required"},
		"originContent": []string{"required"},
		"parseContent":  []string{"required"},
	}

	msg := govalidator.MapData{
		"id": []string{
			"required:文章ID为必填项",
			"exist:文章不存在",
		},
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
