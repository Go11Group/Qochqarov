package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"my_mod/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Client struct {
	http.Client
}

func main() {
	router := gin.Default()
	Client := Client{}

	router.GET("/user", Client.GetUser)
	router.POST("/user", Client.PostUser)
	router.PUT("/user", Client.PutUser)
	router.DELETE("/user", Client.DeleteUser)

	router.Run(":8070")

}

func (cl *Client) GetUser(c *gin.Context) {
	resp, err := http.Get("http://localhost:8090/user")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	c.JSON(http.StatusOK, string(body))
	fmt.Println(string(body))

}

func (cl *Client) PostUser(c *gin.Context) {
	var user model.Users
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	users, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "http://localhost:8090/user", bytes.NewBuffer(users))

	if err != nil {
		panic(err)
	}

	resp, err := cl.Do(req)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&user)
	c.JSON(http.StatusCreated, user)
}

func (cl *Client) PutUser(c *gin.Context) {
	
	var user model.Users
	err := c.BindJSON(&user)
	if err != nil {
		panic(err)
	}
	users, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("PUT", "http://localhost:8090/user", bytes.NewBuffer(users))

	if err != nil {
		panic(err)
	}

	resp, err := cl.Do(req)
	if err != nil {
		panic(err)
	}

	json.NewDecoder(resp.Body).Decode(&user)
	c.JSON(http.StatusCreated, user)
}


func (cl *Client) DeleteUser(c *gin.Context) {
	url := c.Request.URL.Path + "?id=" + c.Query("id")
	fmt.Println(url)
	req, err := http.NewRequest("DELETE", "http://localhost:8090"  + url, nil)

	if err != nil {
		panic(err)
	}

	resp, err := cl.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"Error fatal delete users",
		})
		return
	}
	c.JSON(resp.StatusCode, gin.H{
		"message":"Delete users successully",
	})
}
