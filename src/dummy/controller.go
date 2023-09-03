package dummy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Return simple dummy object
// @Tags dummy
// @Produce json
// @Success 200
// @Param id path string true "Dummy id"
// @Router /dummy/{id} [get]
func findById(c *gin.Context) {
	c.JSON(http.StatusOK, "This is a dummy object with id:"+" "+c.Param("id"))
}

func RegisterRoutes(rg *gin.RouterGroup) {
	dummy := rg.Group("/dummy/")

	dummy.GET(":id", findById)
}
