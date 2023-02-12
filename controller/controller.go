package controller

import (
	"gosearch/db"
	"gosearch/model"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var body struct {
		Name string
	}
	c.Bind(&body)
	u := model.User1{Name: body.Name}
	db.DB.Create(&u)
	c.JSON(200, u)
}

func NgramSearch(c *gin.Context) {
	var body struct {
		Name string
	}
	c.Bind(&body)
	var users []model.User1
	if err := db.DB.Raw("select * from user1 where match name against(?)", body.Name).Scan(&users); err != nil {
		c.Status(400)
		// panic(err)
	}
	c.JSON(200,users)
}

func LikeSearch(c *gin.Context) {
	var body struct {
		Name string
	}
	c.Bind(&body)
	var users []model.User1
	if err := db.DB.Where("name like ?", "%" + body.Name + "%").Find(&users); err != nil {
		c.Status(400)
	}
	c.JSON(200,users)
}
