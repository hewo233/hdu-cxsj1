package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hewo233/hdu-cxsj1/common"
	"github.com/hewo233/hdu-cxsj1/db"
	"github.com/hewo233/hdu-cxsj1/module"
	"github.com/hewo233/hdu-cxsj1/shared/consts"
	"log"
	"net/http"
)

func AddBook(c *gin.Context) {

	email := common.GetEmailFromJWT(c)

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
	var books []module.Book
	db.DB.Table("books").Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"data":  books,
	})
}

func GetBookByID(c *gin.Context) {
	var book module.Book
	bid := c.Param("bid")
	db.DB.Table("books").Where("bid = ?", bid).First(&book)

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"data":  book,
	})
}

func DeleteBookByID(c *gin.Context) {
	bid := c.Param("bid")
	db.DB.Table("books").Where("bid = ?", bid).Delete(&module.Book{})

	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
	})
}

func UpdateBookByID(c *gin.Context) {
	bid := c.Param("bid")

	var book module.Book

	err := c.BindJSON(&book)
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
	db.DB.Table("books").Where("bid = ?", bid).First(oldBook)

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
	c.JSON(http.StatusOK, gin.H{
		"errno": 20000,
		"msg":   "OK",
		"data":  oldBook,
	})
}
