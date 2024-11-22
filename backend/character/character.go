package character

type Character struct {
	Name             string     `json:"name"`
	Race             string     `json:"race"`
	CharacterClass   string     `json:"characterClass"`
	Level            int        `json:"level"`
	Attributes       Attributes `json:"attributes"`
	HitPoints        int        `json:"hitPoints"`
	ExperiencePoints int        `json:"experiencePoints"`
}

type Attributes struct {
	Strength     int `json:"strength"`
	Dexterity    int `json:"dexterity"`
	Constitution int `json:"constitution"`
	Intelligence int `json:"intelligence"`
	Wisdom       int `json:"wisdom"`
	Charisma     int `json:"charisma"`
}

// just adding struct not hooked up yet
type SavingThrows struct {
	Strength     bool `json:"strength"`
	Dexterity    bool `json:"dexterity"`
	Constitution bool `json:"constitution"`
	Intelligence bool `json:"intelligence"`
	Wisdom       bool `json:"wisdom"`
	Charisma     bool `json:"charisma"`
}

var characters = []Character{}
