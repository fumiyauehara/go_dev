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
		result := database.User{}.Find(3)
		context.JSON(http.StatusOK, result)
	})

	_ = r.Run(":9000")
}
