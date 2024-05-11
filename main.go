// main.go
package main

import (
	"log"
	"os"
	"path/filepath"
	"treezoo_backend/api"
	"treezoo_backend/db"
	"treezoo_backend/docs"

	"time"

	"github.com/gin-contrib/cors"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// log設定
	logDir := filepath.Join("C:", "apps", "treezoo", "logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}
	logPath := filepath.Join(logDir, "go_app.log")
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	log.SetOutput(logFile)
	log.Println("Logging started")

	// DB接続
	config := db.LoadConfig("config.yml")
	connection := db.OpenConnection(config)
	defer connection.Close()

	// APIドキュメント化するAPIのベースパスを設定
	docs.SwaggerInfo.BasePath = "/api/v1"

	// Ginライブラリを利用しAPIハンドラを起動
	router := api.SetupRouter(connection)

	// CORS設定
	router.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"*", // すべてのオリジンを許可
		},
		// アクセスを許可したいHTTPメソッド
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	// ルーティング設定
	// /apiディレクトリ配下で定義したAPIを /api/v1 で適宜階層化してエンドポイント化
	v1 := router.Group("/api/v1")
	{
		example := v1.Group("/example")
		{
			example.GET("/ping", api.Ping)
		}
		family := v1.Group("/family")
		{
			family.GET("/relations", api.FetchParentChildRelations(connection))
			family.GET("/animals", api.FetchAnimals(connection))
			family.GET("/animals/relations", api.FetchAnimalsWithRelations(connection))
			family.GET("/children/:parentId", api.FetchChildrenByParentId(connection))
			family.GET("/parents/:childId", api.FetchParentsByChildId(connection))
		}
		picture := v1.Group("/picture")
		{
			picture.GET("/animal/profile/:animalId", api.FetchProfilePictureByAnimalId(connection))
		}
	}

	// Swagger UIのルート設定
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// APIサーバ起動
	router.Run(":8080")
}
