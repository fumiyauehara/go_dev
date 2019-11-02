package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/fumiya-uehara/go_dev/database"
)

func main() {
	r := gin.Default()

	r.GET("/structJSON_test", func(context *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "foo"
		msg.Message = "これは構造体をJSONで返すためのテストです。"
		msg.Number = 1111
		context.JSON(http.StatusOK, msg)
	})

	r.GET("/get_user", func(context *gin.Context) {
		result := database.User{}.FindUser(1)
		context.JSON(http.StatusOK, result)
	})

	r.GET("/store_user", func(context *gin.Context) {
		database.User{FirstName: "Test", LastName: "Taro", Age: 39}.StoreUser()
		context.JSON(http.StatusCreated, "create is OK")
	})

	_ = r.Run(":9000")
}
