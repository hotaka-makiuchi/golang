package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"./models"
)

type event struct {
	ID        int
	Summary   string
	StartDate string
	EndDate   string
	Count     int
}

type databaseModule struct {
	db *gorm.DB
}

func (dbm *databaseModule) open() {
	dbm.db = dbm.gormConnect()
	// テーブル名が複数形出ない場合は指定する
	// dbm.db.SingularTable(true)

	// 実行ログ出力
	dbm.db.LogMode(true)
}

func (dbm *databaseModule) close() {
	dbm.db.Close()
}

func (dbm *databaseModule) gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "P@ssw0rd"
	PROTOCOL := "tcp(127.0.0.1:13306)"
	// DBNAME := "gormsample"
	DBNAME := "dsaf"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func (dbm *databaseModule) where(whereClause ...interface{}) *gorm.DB {
	a := dbm.db
	if whereClause != nil {
		a = dbm.db.Where(whereClause[0], whereClause[1:]...)
	}
	return a
}

func (dbm *databaseModule) first(target interface{}, where ...interface{}) {
	dbm.where(where...).First(target)
	fmt.Printf("%#v\n", target)
}

func (dbm *databaseModule) find(target interface{}, where ...interface{}) {
	dbm.where(where...).Find(target)
	fmt.Printf("%#v\n", target)
}

func (dbm *databaseModule) create(target interface{}) {
	dbm.db.Create(target)
}

func newDb() databaseModule {
	dbm := databaseModule{}
	dbm.open()
	return dbm
}

func main() {
	db := newDb()
	defer db.close()

	// Event構造体のインスタンス化
	e1 := event{}

	// ID = 3 のレコードを１件取得する
	e1.ID = 3
	db.first(&e1)

	// summary = 'hoge4'
	e2 := event{}
	db.first(&e2, "summary=?", "hoge4")

	// 複数件
	e3 := []event{}
	db.find(&e3, "summary like ? and count = ?", "hoge%", 1)

	u1 := models.User{WelbyID: 123, Created: time.Now(), Modified: time.Now()}
	db.create(&u1)

	db.first(&u1)
}
