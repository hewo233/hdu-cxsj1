package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hewo233/hdu-cxsj1/common"
	"github.com/hewo233/hdu-cxsj1/db"
	"github.com/hewo233/hdu-cxsj1/module"
	"github.com/hewo233/hdu-cxsj1/shared/consts"
	"log"
	"net/http"
	"strconv"
)

func AddBook(c *gin.Context) {

	uid := common.GetUIDFromJWT(c)
	if uid == -1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized, uid not found",
		})
		c.Abort()
		return
	}

	var book module.Book
	err := c.ShouldBind(&book)
	if err != nil {
		log.Println("book bind failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error, BindJSON failed",
		})
		c.Abort()
		return
	}

	book.Uid = uid

	file, err := c.FormFile("cover")
	if file != nil && err == nil {
		filePath := consts.BookCoverPath + file.Filename
		err2 := c.SaveUploadedFile(file, filePath)
		if err2 != nil {
			log.Println("cover save failed", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errno": 50000,
				"msg":   "Internal Server Error, SaveUploadedFile failed",
			})
			c.Abort()
			return
		}

		book.CoverFile = filePath
		db.DB.Table("books").Create(&book)

		c.JSON(http.StatusOK, gin.H{
			"errno": 20000,
			"msg":   "OK",
			"data":  book,
		})
	} else {
		if err != nil {
			if err.Error() == "http: no such file" {

				book.CoverFile = consts.DefaultCoverPath

				db.DB.Table("books").Create(&book)

				c.JSON(http.StatusOK, gin.H{
					"errno": 20001,
					"msg":   "OK with default cover",
					"data":  book,
				})

			} else {

				log.Println("FormFile failed", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"errno": 50000,
					"msg":   "Internal Server Error, FormFile failed",
				})

				c.Abort()
				return
			}
		}
	}
}

func ListBook(c *gin.Context) {

	uid := common.GetUIDFromJWT(c)
	if uid == -1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized, uid not found",
		})
		c.Abort()
		return
	}

	user := module.NewUser()
	db.DB.Preload("Books").Where("uid = ?", uid).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"data":  user.Books,
	})
}

func GetBookByID(c *gin.Context) {

	uid := common.GetUIDFromJWT(c)
	if uid == -1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized, uid not found",
		})
		c.Abort()
		return
	}

	var book module.Book

	bidStr := c.Param("bid")
	bid, err := strconv.Atoi(bidStr)
	if err != nil {
		log.Println("bidStr to bid failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error, BidStr to Bid failed",
		})
		c.Abort()
		return
	}

	result := db.DB.Table("books").Where("bid = ? AND uid = ?", bid, uid).First(&book)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"errno": 40400,
			"msg":   "not found, book not found",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"data":  book,
	})
}

func DeleteBookByID(c *gin.Context) {

	uid := common.GetUIDFromJWT(c)
	if uid == -1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized, uid not found",
		})
		c.Abort()
		return
	}

	bidStr := c.Param("bid")
	bid, err := strconv.Atoi(bidStr)
	if err != nil {
		log.Println("bidStr to bid failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error, BidStr to Bid failed",
		})
		c.Abort()
		return
	}

	result := db.DB.Table("books").Where("bid = ? AND uid = ?", bid, uid).Delete(&module.Book{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"errno": 40400,
			"msg":   "not found, book not found",
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
	})
}

func UpdateBookByID(c *gin.Context) {

	uid := common.GetUIDFromJWT(c)
	if uid == -1 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"errno": 40100,
			"msg":   "Unauthorized, uid not found",
		})
		c.Abort()
		return
	}

	bidStr := c.Param("bid")
	bid, err := strconv.Atoi(bidStr)
	if err != nil {
		log.Println("bidStr to bid failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error, BidStr to Bid failed",
		})
		c.Abort()
		return
	}

	var book module.Book

	err = c.ShouldBind(&book)
	if err != nil {
		log.Println("book bind failed", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errno": 50000,
			"msg":   "Internal Server Error, BindJSON failed",
		})
		c.Abort()
		return
	}

	oldBook := module.NewBook()
	result := db.DB.Table("books").Where("bid = ? AND uid = ?", bid, uid).First(oldBook)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"errno": 40400,
			"msg":   "not found, book not found",
		})
		c.Abort()
		return
	}

	if book.Name != "" {
		oldBook.Name = book.Name
	}
	if book.Author != "" {
		oldBook.Author = book.Author
	}
	if book.Publisher != "" {
		oldBook.Publisher = book.Publisher
	}
	if book.Intro != "" {
		oldBook.Intro = book.Intro
	}

	file, err := c.FormFile("cover")
	if file != nil && err == nil {

		filePath := consts.BookCoverPath + file.Filename

		err2 := c.SaveUploadedFile(file, filePath)

		if err2 != nil {
			log.Println("cover save failed", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"errno": 50000,
				"msg":   "Internal Server Error, SaveUploadedFile failed",
			})
			c.Abort()
			return
		}

		oldBook.CoverFile = filePath
	} else {
		if err != nil {
			if err.Error() != "http: no such file" {
				log.Println("FormFile failed", err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"errno": 50000,
					"msg":   "Internal Server Error, FormFile failed",
				})
				c.Abort()
				return
			}
		}
	}

	db.DB.Table("books").Updates(oldBook)

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"data":  oldBook,
	})

}
