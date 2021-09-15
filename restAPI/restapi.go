package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// GET
func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "home page",
	})
}

// POST
func PostHome(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func main() {
	r := gin.Default()
	r.GET("/", GetHome)
	r.POST("/", PostHome)
	r.Run()
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port

}
