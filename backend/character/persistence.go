package character

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func writeCharacterToCSV(c Character) error {
	file, err := os.OpenFile("characters.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	attributesJSON, err := json.Marshal(c.Attributes)
	if err != nil {
		return fmt.Errorf("failed to marshal attributes: %v", err)
	}

	record := []string{
		c.Name,
		c.Race,
		c.CharacterClass,
		fmt.Sprint(c.Level),
		string(attributesJSON),
	}

	if err := writer.Write(record); err != nil {
		return err
	}

	defer writer.Flush()
	return nil

}

var loadCharactersFromCSV = func() error {
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

		var attributes Attributes
		if err := json.Unmarshal([]byte(record[4]), &attributes); err != nil {
			return err
		}

		character := Character{
			Name:           record[0],
			Race:           record[1],
			CharacterClass: record[2],
			Level:          level,
			Attributes:     attributes,
		}

		characters = append(characters, character)
	}

	return nil
}
