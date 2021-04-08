package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//func main() {
//	fmt.Println("climer server start ...")
//
//	timer := time.NewTimer(60 * time.Second)
//	go func() {
//		getHots.Collect()
//		for {
//			timer.Reset(60 * 5 * time.Second) // 这样来复用 timer 和修改执行时间
//			select {
//			case <-timer.C:
//				getHots.Collect()
//			}
//		}
//	}()
//	webPublish.Publish()
//
//}

type Product struct {
	gorm.Model //嵌入常用字段
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "lite.db")
	if err != nil {
		panic(err)
	}
	//关闭数据库连接
	defer db.Close()

	//创建表
	db.AutoMigrate(&Product{})

	product := &Product{Code: "No.002", Price: 520}
	fmt.Println(db.NewRecord(product)) //output: true
	db.Create(product)
	fmt.Println(db.NewRecord(product)) //output: false
}
