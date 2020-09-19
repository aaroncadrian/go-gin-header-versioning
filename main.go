package main

import (
	"github.com/aaroncadrian/go-gin-header-versioning/versioning"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	v := versioning.NewUtil("X-Header")

	r.GET("/ping", v.MapVersions(versioning.VersionMap{
		"2020-09-19": func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		},
	})...)

	err := r.Run()

	if err != nil {
		log.Fatal(err)
	}
}
