package character

import (
	"github.com/gin-gonic/gin"
)

func CharacterRoutes(router *gin.Engine) {
	router.POST("/characters", CreateCharacter)
	router.GET("/characters/:name", GetCharacter)
	router.GET("/characters", ListCharacters)
	router.DELETE("/delete", DeleteCharacters)
}
