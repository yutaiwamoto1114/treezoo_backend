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
// @Success 200 {string} OK
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
			var childId, childName, childSpecies, parentId, parentName, parentSpecies, zooName, zooLocation string

			// rows.Scan()は、SQLクエリの結果から得られる各行のカラムを、指定されたポインタリストに順番に割り当てていく
			// したがって、SQLクエリで指定されたカラムの順番とマッピングの順番は完全に一致している必要がある
			// ORM(ORマッパー)を使わず自力でマッピングしようとするならこうなる
			if err := rows.Scan(&childId, &childName, &childSpecies, &parentId, &parentName, &parentSpecies, &zooName, &zooLocation); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error"})
				return
			}
			result := map[string]interface{}{
				"childId":       childId,
				"childName":     childName,
				"childSpecies":  childSpecies,
				"parentId":      parentId,
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
// @Success 200 {string} OK
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

// @Summary 子取得
// @Schemes
// @Description あるノードについて、その子をすべて取得します。
// @Tags family
// @Accept json
// @Produce json
// @Param parentId path int true "親の動物ID"
// @Success 200 {string} OK
// @Router /family/children/{parentId} [get]
func FetchChildrenByParentId(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		parentId := c.Param("parentId")
		query := `
		SELECT
			a.animal_id AS child_id,
			a.name AS child_name
		FROM
			parent_child_relations r
		JOIN
			animals a ON r.child_id = a.animal_id
		WHERE
			r.parent_id = ?;
        `
		rows, err := db.Query(query, parentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Query error"})
			return
		}
		defer rows.Close()

		var parents []model.AnimalSummary
		for rows.Next() {
			var parent model.AnimalSummary
			if err := rows.Scan(&parent.AnimalId, &parent.AnimalName); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error"})
				return
			}
			parents = append(parents, parent)
		}
		c.JSON(http.StatusOK, parents)
	}
}

// @Summary 親取得
// @Schemes
// @Description あるノードについて、その親をすべて取得します。
// @Tags family
// @Accept json
// @Produce json
// @Param childId path int true "子の動物ID"
// @Success 200 {string} OK
// @Router /family/parents/{childId} [get]
func FetchParentsByChildId(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		childId := c.Param("childId")
		query := `
		SELECT
			a.animal_id AS parent_id,
			a.name AS parent_name
		FROM
			parent_child_relations r
		JOIN
			animals a ON r.parent_id = a.animal_id
		WHERE
			r.child_id = ?;
        `
		rows, err := db.Query(query, childId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Query error"})
			return
		}
		defer rows.Close()

		var children []model.AnimalSummary
		for rows.Next() {
			var child model.AnimalSummary
			if err := rows.Scan(&child.AnimalId, &child.AnimalName); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Scan error"})
				return
			}
			children = append(children, child)
		}
		c.JSON(http.StatusOK, children)
	}
}
