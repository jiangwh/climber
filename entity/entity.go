package entity

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

var db *gorm.DB

var tableName = "hot"

type Hot struct {
	gorm.Model
	Name       string
	Context    string
	DataType   string
	EnableShow int
	Rss        string
}

func init() {
	var err error
	db, err = gorm.Open("sqlite3", "lite.db")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.SetLogger(gorm.Logger{})
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	//defer db.Close()
}

func CreateTable() {
	if !db.HasTable(&Hot{}) {
		db.CreateTable(&Hot{})
		if db.Error != nil {
			panic(db.Error)
		}
	}
}

func (hot *Hot) AddHot() {
	//db.Table(tableName).Create(hot)
	db.AutoMigrate(&Hot{})
	fmt.Println(db.Table(tableName).NewRecord(hot))
	db.Create(hot)

}

func (hot *Hot) RemoveHot() {
	db.Table(tableName).Delete(hot)
}

func (hot *Hot) UpdateHotById() {
	db.Model(&hot).Update(hot)
}

func (hot *Hot) UpdateHotByName(name string) {
	db.Model(hot).Where("name=?", name).Update(hot)
}

func (hot *Hot) QueryHotbyId(id int) *Hot {
	db.Select("id,name,context,dataType,enableShow,rss").
		Where(map[string]interface{}{"id": id}).
		Find(hot)
	return hot
}
