package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hewo233/hdu-cxsj1/db"
	"github.com/hewo233/hdu-cxsj1/module"
	"github.com/hewo233/hdu-cxsj1/shared/consts"
	"github.com/hewo233/hdu-cxsj1/utils/jwt"
	passwd "github.com/hewo233/hdu-cxsj1/utils/password"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	newUser := module.NewUser()
	err := c.BindJSON(newUser)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"errno": 40000,
			"msg":   "Bad Request",
		})
	}

	HashedPassword, err := passwd.HashPassword(newUser.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error",
		})
	}

	newUser.Password = HashedPassword

	db.DB.Table("user").Create(newUser)

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"user":  newUser,
	})

}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := module.NewUser()
	db.DB.Table("user").Where("email = ?", email).First(user)

	err := passwd.CheckHashed(password, user.Password)
	if err != nil {
		log.Println("CheckHashed error: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized at login, password error",
		})
	}

	token, err := jwt.GenerateJWT(user.Name, user.Uid, consts.User)
	if err != nil {
		log.Println("generate JWT token error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"token": token,
	})
}
