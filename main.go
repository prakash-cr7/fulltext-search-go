package main

import (
	"gosearch/controller"
	"gosearch/db"
	"os"

	"github.com/gin-gonic/gin"
)

func init ()  {
	db.InitDb()
}

func main()  {
	os.Setenv("PORT", "8081")
	r := gin.Default()
	r.POST("/createUser", controller.CreateUser)
	r.POST("/ngramSearch", controller.NgramSearch)
	r.POST("/likeSearch", controller.LikeSearch)
	r.Run()
}