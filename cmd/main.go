package main

import (
	"api_consulta_dados/controller"
	"api_consulta_dados/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database := db.Connectiondb()
	defer database.Close()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "Hello World",
		})
	})

	r.GET("/cadsus/cpf", func(c *gin.Context) {
		controller.ConsultCpf(c, database)
	})
	r.GET("/cadsus/nome", func(c *gin.Context) {
		controller.ConsultNome(c, database)
	})
	r.GET("/cadsus/celular", func(c *gin.Context) {
		controller.ConsultCelular(c, database)
	})

	r.Run(":8080")
}
