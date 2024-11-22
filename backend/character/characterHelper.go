package character

func calculateHitPoints(c Character) int {

	var hitDie int
	switch c.CharacterClass {
	case "Fighter":
		hitDie = 10
	case "Barbarian":
		hitDie = 12
	case "Wizard":
		hitDie = 6
	case "Rogue":
		hitDie = 8
	case "Cleric":
		hitDie = 8
	case "Druid":
		hitDie = 8
	case "Monk":
		hitDie = 8
	case "Paladin":
		hitDie = 10
	case "Sorcerer":
		hitDie = 6
	case "Bard":
		hitDie = 8
	case "Warlock":
		hitDie = 8
	case "Ranger":
		hitDie = 10
	default:
		hitDie = 8
	}

	hp := hitDie + (c.Attributes.Constitution-10)/2

	//TODO increase hp based on level
	return hp
}
