package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("static/*.html")

	data := "Hello Go/Gin!"

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"data": data})
	})
	r.Run(":8000")
}