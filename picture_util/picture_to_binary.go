// picture_to_binary.go
package main

import (
	"encoding/base64"
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
	// エンコード前のデータサイズ
	fmt.Println("Original Data Size:", len(data), "bytes")

	// Base64エンコード
	encodedData := base64.StdEncoding.EncodeToString(data)
	// エンコード後のデータサイズ
	fmt.Println("Encoded Data Size:", len(encodedData), "bytes")

	// DBにバイナリデータ挿入
	stmt, err := con.Prepare("insert into pictures (picture_data) values (?)")
	if err != nil {
		log.Fatal(err)
	}
	// stmt.Exec(encodedData)
	stmt.Exec(data)

	fmt.Println("画像がDBに保存されました。")
}
