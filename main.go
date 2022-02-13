package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nome string
}

func main() {

	databese_dsn := "root:root@/?parseTime=true"

	db, err := gorm.Open(mysql.Open(databese_dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Conn db failed")
	}
	db.Exec("create database if not exists gogin")
	db.Exec("use gogin")

	db.AutoMigrate(&User{})
	// db.Create(&User{Nome: "User 1"})

	page_detail := struct {
		title string
		desc  string
	}{title: "Goglang - Gin", desc: "CRUD"}

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	// INDEX
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": page_detail.title, "desc": page_detail.desc})
	})

	// GET ALL USERS
	router.GET("/users", func(c *gin.Context) {
		users := []User{}
		db.Table("users").Find(&users)
		log.Printf("Users %v", users)
		c.HTML(http.StatusOK, "users.html", gin.H{"title": page_detail.title, "desc": page_detail.desc, "users": &users})
	})

	router.Run(":8000")
}
