package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/syutingsong/regions/db"
	"net/http"
)

func Provinces(c *gin.Context) {
	if c.Param("short") == "short" {
		c.JSON(http.StatusOK, db.ShortProvinces)
	} else {
		c.JSON(http.StatusOK, db.Provinces)
	}
}
