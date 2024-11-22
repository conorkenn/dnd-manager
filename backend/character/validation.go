package character

import (
	"encoding/csv"
	"fmt"
	"os"
)

var ValidClasses = []string{"Barbarian", "Bard", "Cleric", "Druid", "Fighter", "Monk", "Paladin", "Ranger", "Rogue", "Sorcerer", "Warlock", "Wizard"}
var ValidRaces = []string{"Human", "Elf", "Dwarf", "Halfling", "Gnome", "Half-Elf", "Half-Orc", "Tiefling", "Dragonborn"}

func validateCharacter(character Character) error {

	taken, err := isNameTaken(character.Name)
	if err != nil {
		return fmt.Errorf("failed to check csv %v", err)
	}

	if taken {
		return fmt.Errorf("name is already taken")
	}

	if !isValidRace(character.Race) {
		return fmt.Errorf("not a valid race")
	}

	if !isValidClass(character.CharacterClass) {
		return fmt.Errorf("not a valid class")
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

func isValidRace(race string) bool {
	for _, r := range ValidRaces {
		if r == race {
			return true
		}
	}
	return false
}

func isValidClass(class string) bool {
	for _, c := range ValidClasses {
		if c == class {
			return true
		}
	}
	return false
}
