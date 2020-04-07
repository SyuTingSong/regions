package main

import (
	"github.com/gin-gonic/gin"
	"github.com/integrii/flaggy"
	"github.com/syutingsong/regions/actions"
	"github.com/syutingsong/regions/db"
	"log"
	"os"
)

func main() {
	db.InitShort()

	var bind, ok = os.LookupEnv("GIN_BIND")
	if !ok {
		bind = "[::]:8080"
	}
	flaggy.String(&bind, "b", "bind", "binding ip:port")
	flaggy.Parse()

	r := gin.Default()
	r.GET("/regions/:short/provinces", actions.Provinces)
	r.GET("/regions/:short/cities", actions.Cities)
	r.GET("/regions/:short/districts", actions.District)
	if err := r.Run(bind); err != nil {
		log.Fatalln("listen error", err)
	}
}
