package main

import (
	"dnd-manager/character"
	"dnd-manager/dice"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	router.POST("/characters", character.CreateCharacter)
	router.GET("/characters", character.ListCharacters)
	router.GET("/characters/:name", character.GetCharacter)
	router.DELETE("/delete", character.DeleteCharacters)
	router.POST("/roll", dice.Roll)

	router.Run(":8080")
}
