package boot

import (
	"DJ-Blog/internal/model"
	"DJ-Blog/internal/service"
	"DJ-Blog/pkg/database"
)

func Initialize() {

	// 数据库自动建表
	// gorm默认建表规则如下
	// 1. 表名为结构体名的蛇形复数形式
	// 2. 字段名为结构体字段名的蛇形形式
	// 这里的数据库引擎是MYISAM，有更高的搜索效率，但是不支持事务
	database.DB.Set("gorm:table_options", "ENGINE=MYISAM").
		AutoMigrate(&service.User{},
			&service.Article{},
			&model.TagModel{},
			&model.CategoryModel{})
}
