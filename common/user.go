package common

import (
	"github.com/gin-gonic/gin"
	"github.com/hewo233/hdu-cxsj1/db"
	"github.com/hewo233/hdu-cxsj1/module"
	"log"
	"strconv"
)

func GetEmailFromJWT(c *gin.Context) string {
	email := c.GetString("userEmail")

	user := module.NewUser()

	result := db.DB.Table("users").First(&user, "email = ?", email)
	if result.RowsAffected == 0 {
		return ""
	} else {
		return email
	}
}

func GetUIDFromJWT(c *gin.Context) int {
	uidString := c.GetString("uid")

	uid, err := strconv.Atoi(uidString)
	if err != nil {
		log.Println("uid parse failed", err)
		return -1
	}

	user := module.NewUser()

	result := db.DB.Table("users").First(&user, "uid = ?", uid)
	if result.RowsAffected == 0 {
		log.Println("uid not found")
		return -1
	} else {
		return user.Uid
	}

}
