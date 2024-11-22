package character

var xpBreakPoints = []int{0, 300, 900, 2700, 6500, 14000, 23000, 34000, 48000, 64000, 85000, 100000, 120000, 140000, 165000, 195000, 225000, 265000, 305000, 355000}

func applyXpGains(c *Character, xp int) {
	c.ExperiencePoints += xp
	for i, breakPoint := range xpBreakPoints {
		if c.ExperiencePoints < breakPoint {
			c.Level = i
			return
		}
	}

	c.Level = len(xpBreakPoints)
}

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
