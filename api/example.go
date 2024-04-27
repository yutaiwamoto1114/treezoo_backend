// @BasePath /api/v1
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary ping疎通確認
// @Schemes
// @Description 無条件に"ping OK!"と返します。
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} ping
// @Router /example/ping [get]
func Ping(g *gin.Context) {
	g.JSON(http.StatusOK, "ping OK!")
}
