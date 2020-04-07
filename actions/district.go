package actions

import (
	"github.com/gin-gonic/gin"
	"github.com/syutingsong/regions/db"
	"net/http"
)

func District(c *gin.Context) {
	city, ok := c.GetQuery("city")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "city is required",
		})
		return
	}
	var districts []string
	if c.Param("short") == "short" {
		districts, ok = db.ShortCities[city]
	} else {
		districts, ok = db.Cities[city]
	}
	if !ok {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "city not found",
		})
		return
	}

	c.JSON(http.StatusOK, districts)
}
