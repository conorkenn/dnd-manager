package character

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Character struct {
	Name           string     `json:"name"`
	Race           string     `json:"race"`
	CharacterClass string     `json:"characterClass"`
	Level          int        `json:"level"`
	Attributes     Attributes `json:"attributes"`
}

type Attributes struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
}

var characters = []Character{}

func CreateCharacter(c *gin.Context) {
	var newCharacter Character

	if err := c.BindJSON(&newCharacter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taken, err := isNameTaken(newCharacter.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check csv"})
	}
	if taken {
		c.JSON(http.StatusConflict, gin.H{"message": "character name is already taken"})
		return
	}
	characters = append(characters, newCharacter)

	if err := writeCharacterToCSV(newCharacter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save character"})
	}
	c.JSON(http.StatusOK, gin.H{
		"character": newCharacter,
		"message":   "successfully created character",
	})

}

func ListCharacters(c *gin.Context) {

	if err := loadCharactersFromCSV(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to load characters"})
	}
	c.JSON(http.StatusOK, gin.H{
		"characters": characters,
		"message":    "successfully listed characters"})
}

func DeleteCharacters(c *gin.Context) {
	if err := os.Remove("characters.csv"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete characters"})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully delete characters"})
}

func writeCharacterToCSV(c Character) error {
	file, err := os.OpenFile("characters.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		c.Name,
		c.Race,
		c.CharacterClass,
		fmt.Sprint(c.Level),
		fmt.Sprint(c.Attributes),
	}

	if err := writer.Write(record); err != nil {
		return err
	}

	return nil

}

func loadCharactersFromCSV() error {
	file, err := os.Open("characters.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	characters = []Character{}

	for _, record := range records {
		level, err := strconv.Atoi(record[3])
		if err != nil {
			return err
		}

		character := Character{
			Name:           record[0],
			Race:           record[1],
			CharacterClass: record[2],
			Level:          level,
		}

		characters = append(characters, character)
	}

	return nil
}

func isNameTaken(name string) (bool, error) {
	file, err := os.Open("characters.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return false, err
	}

	for _, record := range records {
		if len(record) > 0 && record[0] == name {
			return true, nil
		}
	}
	return false, nil
}
