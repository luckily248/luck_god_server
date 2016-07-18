package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luckily248/luck_god_server/controller"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/**/*")
	r.Static("/static", "./static")
	r.GET("/event/index", controller.EvenCIndex)
	r.Run()
}
