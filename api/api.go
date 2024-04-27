package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
func ParentChildRelations(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// view_parent_child_relations ビューから全件取得を実行
		rows, err := db.Query("SELECT * FROM view_parent_child_relations")
		if err != nil {
			log.Fatal("Query error:", err)
		}
		defer rows.Close()

		results := make([]map[string]interface{}, 0)
		for rows.Next() {
			var childName, childSpecies, parentName, parentSpecies, zooName, zooLocation string
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
