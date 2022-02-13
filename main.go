package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	page_detail := struct {
		title string
		desc  string
	}{title: "Goglang - Gin", desc: "CRUD"}

	router := gin.Default()

	router.LoadHTMLGlob("templates/**/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": page_detail.title, "desc": page_detail.desc})
	})

	router.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users.html", gin.H{"title": page_detail.title, "desc": page_detail.desc, "test": "a"})
	})

	router.Run(":8000")
}
