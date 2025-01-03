package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hewo233/hdu-cxsj1/common"
	"github.com/hewo233/hdu-cxsj1/db"
	"github.com/hewo233/hdu-cxsj1/module"
	"github.com/hewo233/hdu-cxsj1/shared/consts"
	"github.com/hewo233/hdu-cxsj1/utils/jwt"
	passwd "github.com/hewo233/hdu-cxsj1/utils/password"
	"log"
	"net/http"
	"strconv"
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
		c.Abort()
		return
	}

	tryUser := db.DB.Table("users").Where("email = ?", newUser.Email).First(&module.User{})
	if tryUser.RowsAffected > 0 {
		log.Println("email exists")
		c.JSON(http.StatusBadRequest, gin.H{
			"errno": 40000,
			"msg":   "Bad Request, email exists",
		})
		c.Abort()
		return
	}
	tryUser = db.DB.Table("users").Where("name = ?", newUser.Name).First(&module.User{})
	if tryUser.RowsAffected > 0 {
		log.Println("name exists")
		c.JSON(http.StatusBadRequest, gin.H{
			"errno": 40000,
			"msg":   "Bad Request, name exists",
		})
		c.Abort()
		return
	}

	HashedPassword, err := passwd.HashPassword(newUser.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error",
		})
		c.Abort()
		return
	}

	newUser.Password = HashedPassword

	db.DB.Table("users").Create(newUser)

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"user":  newUser,
	})

}

func Login(c *gin.Context) {

	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var loginReq LoginRequest
	err := c.BindJSON(&loginReq)

	user := module.NewUser()
	db.DB.Table("users").Where("email = ?", loginReq.Email).First(user)
	if user.Name == "" {
		log.Println("user not found")
		c.JSON(http.StatusNotFound, gin.H{
			"errno": 40400,
			"msg":   "Not Found, user not found",
		})
		c.Abort()
		return
	}

	err = passwd.CheckHashed(loginReq.Password, user.Password)
	if err != nil {
		log.Println("CheckHashed error: ", err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized at login, password error",
		})
		c.Abort()
		return
	}

	token, err := jwt.GenerateJWT(user.Email, user.Uid, consts.User)
	if err != nil {
		log.Println("generate JWT token error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error",
		})
		c.Abort()
		return
	}

	type LoginUserResponse struct {
		Uid    int    `json:"uid"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Gender string `json:"gender"`
	}
	userRep := LoginUserResponse{
		Uid:    user.Uid,
		Name:   user.Name,
		Email:  user.Email,
		Gender: user.Gender,
	}

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"token": token,
		"user":  userRep,
	})
}

func GetUserInfoByID(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		log.Println("uidStr to uid failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error, uidStr to uid failed",
		})
		c.Abort()
		return
	}

	jwtUid := common.GetUIDFromJWT(c)
	if jwtUid == -1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized, uid not found",
		})
		c.Abort()
		return
	}

	//log.Printf("GetUserInfoByID: uid: %s jwtID: %s \n", uid, jwtID)

	if jwtUid != uid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized, uid not match",
		})
		c.Abort()
		return
	}

	user := module.NewUser()
	db.DB.Table("users").Where("uid = ?", uid).First(user)

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"user":  user,
	})
}

func UpdateUserInfoByID(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		log.Println("uidStr to uid failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error, uidStr to uid failed",
		})
		c.Abort()
		return
	}

	JWTUid := common.GetUIDFromJWT(c)
	if JWTUid == -1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized, uid not found",
		})
		c.Abort()
		return
	}

	if JWTUid != uid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized, uid not match",
		})
		c.Abort()
		return
	}

	newUser := module.NewUser()
	err = c.BindJSON(newUser)
	if err != nil {
		log.Println("Update user BindJSON failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error, BindJSON failed",
		})
		c.Abort()
		return
	}

	oldUser := module.NewUser()
	db.DB.Table("users").Where("uid = ?", uid).First(oldUser)

	if newUser.Name != "" {
		tyrUser := module.NewUser()
		result := db.DB.Table("users").Where("name = ?", newUser.Name).First(tyrUser)
		if result.RowsAffected > 0 && tyrUser.Uid != uid {
			log.Println("name exists")
			c.JSON(http.StatusBadRequest, gin.H{
				"errno": 40000,
				"msg":   "Bad Request, name exists",
			})
			c.Abort()
			return
		}
		oldUser.Name = newUser.Name
	}
	if newUser.Gender != "" {
		oldUser.Gender = newUser.Gender
	}
	if newUser.Password != "" {
		HashedPassword, err := passwd.HashPassword(newUser.Password)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errno": 50000,
				"msg":   "Internal Server Error",
			})
			c.Abort()
			return
		}
		oldUser.Password = HashedPassword
	}

	db.DB.Table("users").Where("uid = ?", uid).Updates(oldUser)

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"user":  oldUser,
	})
}
