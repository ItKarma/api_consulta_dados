package controller

import "github.com/gin-gonic/gin"

func ConsultCpf(c *gin.Context) {
	c.JSON(200, gin.H{
		"Message": "Controller function is valid ",
	})
}
