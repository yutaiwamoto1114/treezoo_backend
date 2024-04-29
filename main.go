// main.go
package main

import (
	"treezoo_backend/api"
	"treezoo_backend/db"
	"treezoo_backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// DB接続
	config := db.LoadConfig("config.yml")
	database := db.OpenConnection(config)
	defer database.Close()

	// APIドキュメント化するAPIのベースパスを設定
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Ginライブラリを利用しAPIハンドラを起動
	router := api.SetupRouter(database)

	// /apiディレクトリ配下で定義したAPIを /api/v1 で適宜階層化してエンドポイント化
	v1 := router.Group("/api/v1")
	{
		example := v1.Group("/example")
		{
			example.GET("/ping", api.Ping)
		}
		family := v1.Group("/family")
		{
			family.GET("/relations", api.FetchParentChildRelations(database))
			family.GET("/animals", api.FetchAnimals(database))
		}
	}

	// Swagger UIのルート設定
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// APIサーバ起動
	router.Run(":8080")
}
