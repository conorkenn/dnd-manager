package character

type Character struct {
	Name           string     `json:"name"`
	Race           string     `json:"race"`
	CharacterClass string     `json:"characterClass"`
	Level          int        `json:"level"`
	Attributes     Attributes `json:"attributes"`
	HitPoints      int        `json:"hitpoints"`
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
