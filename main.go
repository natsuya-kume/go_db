package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	// db接続
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 関数呼び出し
	addUserData(db, "Vaundy")
	// delete(db, 3)
	// findLike(db, "SaucyDog")
	// update(db, 4, "TOCCHI")

	// データベースからデータ取得
	result := []*Users{}
	error := db.Find(&result).Error

	if error != nil {
		fmt.Println(error)
	}

	result = []*Users{}
	db.Find(&result)
	for _, user := range result {
		fmt.Println(user)
	}

}

//.envを呼び出します。
func loadEnv() {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

}

// データ追加関数
func addUserData(db *gorm.DB, name string) {
	error := db.Create(&Users{
		Name:     name,
		Age:      18,
		Address:  "東京都千代田区",
		UpdateAt: getDate(),
	}).Error
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("データ追加したよ！")
	}
}

// データ削除関数
func delete(db *gorm.DB, id int) {
	error := db.Where("id = ?", id).Delete(Users{}).Error
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("deleteしたよ")
	}
}

func findLike(db *gorm.DB, keyword string) {
	result := []*Users{}
	error := db.Where("name LIKE ?", "%"+keyword+"%").Find(&result).Error
	if error != nil || len(result) == 0 {
		return
	}
	for _, user := range result {
		fmt.Println(user.Name)
		// fmt.Println("探しているデータが見つかったよ！")
	}
}

// アップデート関数
func update(db *gorm.DB, id int, name string) {
	result := []*Users{}
	db.Find(&result)
	for _, user := range result {
		fmt.Println(string(user.ID) + "_" + user.Name)
	}
	fmt.Println("update")
	// Modelに構造体の配列をいれる
	error := db.Model(Users{}).Where("id = ?", id).Update(&Users{
		Name:     name,
		UpdateAt: getDate(),
	}).Error
	//// UPDATE users SET name='ゴン太', update_at={現在日時};

	if error != nil {
		fmt.Println(error)
	}

	result = []*Users{}
	db.Find(&result)
	for _, user := range result {
		fmt.Println(string(user.ID) + "_" + user.Name)
		fmt.Println("アップデート完了！")
	}
}

// 日付取得関数
func getDate() string {
	const layout = "2006-01-02 15:04:05"
	now := time.Now()
	return now.Format(layout)
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {

	dbms := os.Getenv("DBMS")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	protocol := os.Getenv("PROTOCOL")
	dbname := os.Getenv("DBNAME")

	DBMS := dbms
	USER := user
	PASS := password
	PROTOCOL := protocol
	DBNAME := dbname

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

// 構造体の定義
type Users struct {
	ID       int
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	UpdateAt string `json:"updateAt" sql:"not null;type:date"`
}
