package dao

import (
	"IM/models"
	"fmt"
	"os"
)

func migration() {
	//自动迁移模式
	//AutoMigrate函数是把代码映射到数据库中
	//AddForeignKey函数是添加外键
	// 	举例子

	// DB.Model(&Cart{}).AddForeignKey("product_id","Product(id)","CASCADE","CASCADE")
	// 这个函数就是在Cart表中把product_id字段关联到product的id中，后面两个就是update和delete的时候进行级联更新或是级联删除。

	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&models.User{},
			&models.UserRoom{},
			&models.Room{},
			&models.Message{},
		)
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
