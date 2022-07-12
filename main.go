package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/Duelana-Team/duelana-backend-v1/config"
	"github.com/Duelana-Team/duelana-backend-v1/controllers"
	"github.com/Duelana-Team/duelana-backend-v1/db"
)

func main() {

	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	r := gin.Default()
	h := db.Init(config.DBUrl)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  config.Port,
			"dburl": config.DBUrl,
		})
	})

	controllers.RegisterRoutes(r, h)
	r.Run(config.Port)
}
