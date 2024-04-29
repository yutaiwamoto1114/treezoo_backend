package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"treezoo_backend/model"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	return router
}

// @Summary 親子関係全件取得
// @Schemes
// @Description すべての親子関係を取得します。
// @Tags family
// @Accept json
// @Produce json
// @Success 200 {string} ping
// @Router /family/relations [get]
func FetchParentChildRelations(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// view_parent_child_relations ビューから全件取得を実行
		rows, err := db.Query("SELECT * FROM view_parent_child_relations")
		if err != nil {
			log.Fatal("Query error:", err)
		}
		defer rows.Close()

		results := make([]map[string]interface{}, 0)
		for rows.Next() {
			// 結果をマッピングする変数群 (Modelクラスを作っておくのが本来のやり方)
			var childName, childSpecies, parentName, parentSpecies, zooName, zooLocation string

			// rows.Scan()は、SQLクエリの結果から得られる各行のカラムを、指定されたポインタリストに順番に割り当てていく
			// したがって、SQLクエリで指定されたカラムの順番とマッピングの順番は完全に一致している必要がある
			// ORM(ORマッパー)を使わず自力でマッピングしようとするならこうなる
			if err := rows.Scan(&childName, &childSpecies, &parentName, &parentSpecies, &zooName, &zooLocation); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error"})
				return
			}
			result := map[string]interface{}{
				"childName":     childName,
				"childSpecies":  childSpecies,
				"parentName":    parentName,
				"parentSpecies": parentSpecies,
				"zooName":       zooName,
				"zooLocation":   zooLocation,
			}
			results = append(results, result)
		}

		// 結果を返却
		c.JSON(http.StatusOK, results)
	}
}

// @Summary 動物全件取得
// @Schemes
// @Description すべての動物を取得します。
// @Tags family
// @Accept json
// @Produce json
// @Success 200 {string} animals
// @Router /family/animals [get]
func FetchAnimals(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// view_animals ビューから全件取得を実行
		rows, err := db.Query(`select * from view_animals`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Query error: " + err.Error()})
			return
		}
		defer rows.Close()

		// 結果配列
		var animals []model.Animal

		for rows.Next() {
			// 1レコードをマッピングするModelクラスのインスタンス
			var animal model.Animal

			// Modelクラスのインスタンスフィールドに値を順番にマッピングする
			// 誕生日(Birthday)など、NULLが許容されるフィールドは明示的な制御が必要
			if err := rows.Scan(
				&animal.AnimalID, &animal.AnimalName, &animal.Species,
				&animal.Birthday, &animal.Age, &animal.Gender,
				&animal.BirthZooID, &animal.BirthZooName, &animal.BirthZooLocation,
				&animal.CurrentZooID, &animal.CurrentZooName, &animal.CurrentZooLocation,
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error: " + err.Error()})
				return
			}
			animals = append(animals, animal)
		}

		// 結果を返却
		c.JSON(http.StatusOK, animals)
	}
}
