package common

import (
	"github.com/gin-gonic/gin"
	"github.com/hewo233/hdu-cxsj1/db"
	"github.com/hewo233/hdu-cxsj1/module"
)

func GetEmailFromJWT(c *gin.Context) string {
	email := c.GetString("userEmail")

	user := module.NewUser()

	result := db.DB.Table("user").First(&user, "email = ?", email)
	if result.RowsAffected == 0 {
		return ""
	} else {
		return email
	}
}
