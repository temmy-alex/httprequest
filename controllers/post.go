package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"http-request/middleware"
	"http-request/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	var result []models.Post

	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	fmt.Println(res.Body)
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	defer res.Body.Close()

	c.JSON(http.StatusOK, gin.H{
		"Data": result,
	})
}

func GetPostById(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	var result models.Post
	var id = c.Param("id")
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	fmt.Println(res.Body)
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	defer res.Body.Close()

	c.JSON(http.StatusOK, gin.H{
		"Data": result,
	})
}

func CreatePost(c *gin.Context) {
	if !middleware.Auth(c) {
		return
	}

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	reqJson, err := json.Marshal(post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	json.Unmarshal(body, &post)

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}
