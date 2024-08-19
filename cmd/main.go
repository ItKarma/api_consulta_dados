package main

import (
	"api_consulta_dados/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "Hello World",
		})
	})

	r.GET("/cpf", controller.ConsultCpf)

	r.Run(":8080")
}
