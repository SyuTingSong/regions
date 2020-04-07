package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/syutingsong/regions/db"
	"net/http"
)

func Cities(c *gin.Context) {
	province, ok := c.GetQuery("province")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "province is required",
		})
		return
	}

	var cities []string
	if c.Param("short") == "short" {
		cities, ok = db.ShortProvinces[province]
	} else {
		cities, ok = db.Provinces[province]
	}
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "no such province",
		})
		return
	}

	c.JSON(http.StatusOK, cities)
}
