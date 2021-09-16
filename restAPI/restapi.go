package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET , "/"
func GetHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "home page",
	})
}

/*
Querystring parameters
*/

// GET - query , "/query"
func GetQueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

// GET - query , "/welcome"
func GetQueryStrings2(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

/*
Parameter in path
*/

// GET - parameter , "/path"
func GetPathParameter(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

// GET - parameter , "/user"
func GetPathParameter2(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

/*
POST
*/

// POST , "/"
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

/*
Multipart/Urlencoded Form
*/

func PostMultipart(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "annoymous")

	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

/*
query + post form
*/

func queryPostForm(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}

/*
Map as querystring or postform parameters
*/

func PostMap(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")

	c.String(http.StatusOK, "ids: %v; names: %v", ids, names)
}

func main() {
	r := gin.Default()
	r.GET("/", GetHome)
	r.GET("/query", GetQueryStrings)            // /query?name=jm&age=29
	r.GET("/welcome", GetQueryStrings2)         // /welcome?firstname=Jongmin&lastname=Han
	r.GET("/path/:name/:age", GetPathParameter) // /path/jm/29
	r.GET("/user/:name", GetPathParameter2)     // /user/jm
	r.POST("/", PostHome)
	r.POST("/form_post", PostMultipart) // /form_post
	r.POST("/post", queryPostForm)
	r.POST("/post_map", PostMap)

	r.Run()
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port

}
