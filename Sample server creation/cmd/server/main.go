package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	PORT := "5050"
	r := gin.Default()

	r.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"Message":fmt.Sprintf("Server is running on the port %s", PORT),
		})
	})

	r.GET("/ping",func(c * gin.Context){
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})


	if err := r.Run((fmt.Sprintf(":%s",PORT))) ; err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}