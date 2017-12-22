package main

import (
	_"encoding/json"
	_"fmt"
	_"html"
	_"io/ioutil"
	_"log"
	"net/http"
	_"os"

	"github.com/gin-gonic/gin"
	m"github.com/stn117/tmBGOt/models"
)

func setupError(r* gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "errors/404.tmpl.html", nil)
	})
}


func update(c *gin.Context) {
	var json m.Message
	if err := c.ShouldBindJSON(&json); err == nil {
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(200, gin.H{
		"from":  json.From,
		"message": json.Text,
		"id":    json.MessageID,
	})
}


func SetupUrls(r * gin.Engine) {

	setupError(r)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	r.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	r.POST("/tmbgot", update)

	//r.GET("/login", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	//})


	//auth := r.Group("/v1")
	//auth.Use()
	//{
	//	auth.POST("/stats", loginEndpoint)
	//	auth.GET("/users", submitEndpoint)
	//	auth.GET("/motherfucker", readEndpoint)
	//}

}

