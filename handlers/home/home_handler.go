package handlers

import "github.com/gin-gonic/gin"

func GetHomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "API Home berjalan",
	})

}
