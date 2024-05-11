// picture_to_binary.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"treezoo_backend/db"
)

func main() {

	// DB接続
	config := db.LoadConfig("config.yml")
	con := db.OpenConnection(config)
	defer con.Close()

	// 画像ファイルの読み込み
	data, err := ioutil.ReadFile(`C:\Users\yutai\OneDrive\画像\カメラ ロール\TreeZoo\Hippo\animal_hippo.png`)
	if err != nil {
		log.Fatal(err)
	}

	// DBにバイナリデータ挿入
	stmt, err := con.Prepare("insert into pictures (picture_data) values (?)")
	if err != nil {
		log.Fatal(err)
	}
	stmt.Exec(data)

	fmt.Println("画像がDBに保存されました。")
}
