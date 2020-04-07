package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/syutingsong/regions/db"
	"net/http"
)

func Provinces(c *gin.Context) {
	var pMap map[string][]string
	if c.Param("short") == "short" {
		pMap = db.ShortProvinces
	} else {
		pMap = db.Provinces
	}
	var provinces []string
	for province := range pMap {
		provinces = append(provinces, province)
	}
	c.JSON(http.StatusOK, provinces)
}
