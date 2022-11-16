package model

type CategoryModel struct {
	Id   uint32 `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name"`
}

func (CategoryModel) TableName() string {
	return "article_category"
}
