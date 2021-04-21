package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "top.html", gin.H{
			"message": "My name is Hoge",
		})
	})
	router.Run()
}
