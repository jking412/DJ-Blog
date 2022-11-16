package model

type CategoryModel struct {
	Id   uint64 `gorm:"column:id;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;type:varchar(100);not null"`
}

func (CategoryModel) TableName() string {
	return "category"
}
