package service

import (
	"TestImport/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

func AddPerson(person models.Person) {
	db, err := gorm.Open("mssql", "sqlserver://sa:liyang@localhost:1433?database=gotest")
	if err != nil {
		fmt.Println("连接数据库失败")
		panic("连接数据库失败")
	}
	db.LogMode(true)
	defer db.Close()
	// 自动迁移模式
	//db.AutoMigrate(&Product{})

	// 添加数据表，返回受影响行数
	result := db.Create(&person).RowsAffected

	if (err != nil) {
		fmt.Println(err)
	}

	fmt.Printf("结果:%v", result)
}
