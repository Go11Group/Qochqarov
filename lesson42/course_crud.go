package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"my_mod/model"
	"github.com/gin-gonic/gin"
)

func (cl *Client) GetCourse(c *gin.Context) {
	resp, err := http.Get("http://localhost:8090/user")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	c.JSON(http.StatusOK, string(body))
	fmt.Println(string(body))

}

func (cl *Client) PostCourse(c *gin.Context) {
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

func (cl *Client) PutCourse(c *gin.Context) {

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

func (cl *Client) DeleteCourse(c *gin.Context) {
	url := c.Request.URL.Path
	var user model.Users

	req, err := http.NewRequest("PUT", "http://localhost:8090"+url, nil)

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
