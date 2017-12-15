package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func setupError(r* gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "errors/404.tmpl.html", nil)
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

