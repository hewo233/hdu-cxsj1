package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hewo233/hdu-cxsj1/Init"
	"github.com/hewo233/hdu-cxsj1/route"
)

func main() {
	Init.Init()

	r := gin.Default()

	route.InitRoute(r)

	r.Run(":8080")
}
