package route

import (
	"github.com/gin-gonic/gin"
	"github.com/hewo233/hdu-cxsj1/handler"
	"github.com/hewo233/hdu-cxsj1/middleware"
)

func InitRoute(r *gin.Engine) {
	r.GET("/ping", handler.Ping)
	auth := r.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}
	user := r.Group("/user")
	user.Use(middleware.JWTAuth("user"))
	{
		user.GET("/:uid", handler.GetUserInfoByID)
	}

	book := r.Group("/book")
	book.Use(middleware.JWTAuth("user"))
	{
		book.POST("/add", handler.AddBook)
		book.GET("/list", handler.ListBook)
		book.GET("/:bid", handler.GetBookByID)
		book.PUT("/:bid", handler.UpdateBookByID)
		book.DELETE("/:bid", handler.DeleteBookByID)
	}
}
