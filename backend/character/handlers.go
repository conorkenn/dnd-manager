package character

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateCharacter(c *gin.Context) {
	var newCharacter Character

	if err := c.BindJSON(&newCharacter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Character data received:", newCharacter) // Log the character data

	if err := validateCharacter(newCharacter); err != nil {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		return
	}

	characters = append(characters, newCharacter)

	if err := writeCharacterToCSV(newCharacter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save character"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"character": newCharacter,
		"message":   "successfully created character",
	})
}

func GetCharacter(c *gin.Context) {
	name := c.Param("name")

	if err := loadCharactersFromCSV(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load characters"})
		return
	}

	for _, character := range characters {
		if character.Name == name {
			c.JSON(http.StatusOK, character)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "character not found"})
}

func ListCharacters(c *gin.Context) {

	if err := loadCharactersFromCSV(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load characters"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"characters": characters,
		"message":    "successfully listed characters"})
}

func DeleteCharacters(c *gin.Context) {
	if err := os.Remove("characters.csv"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete characters"})
		return
	}
	characters = []Character{}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete characters"})
}
