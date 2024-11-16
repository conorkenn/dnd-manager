package character

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Character struct {
	Name  string `json:"name"`
	Race  string `json:"race"`
	Class string `json:"class"`
	Level int    `json:"level"`
}

var characters = []Character{}

func CreateCharacter(c *gin.Context) {
	var newCharacter Character

	if err := c.BindJSON(&newCharacter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	characters = append(characters, newCharacter)
	c.JSON(http.StatusOK, newCharacter)
}

func ListCharacters(c *gin.Context) {
	c.JSON(http.StatusOK, characters)
}
