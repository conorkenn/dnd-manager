package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	/*router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "dnd manager",
		})
	})
	*/

	router.POST("/characters", createCharacter)
	router.GET("/characters", listCharacters)

	router.Run(":8080")
}
