package model

type TagModel struct {
	Id   uint32 `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name"`
}

func (TagModel) TableName() string {
	return "article_tag"
}
