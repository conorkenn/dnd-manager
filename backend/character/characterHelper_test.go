package character

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterHelper(t *testing.T) {
	testCharacter := Character{}
	testCharacter.Attributes.Constitution = 7

	t.Run("Calculate Hit Points", func(t *testing.T) {
		testCharacter.CharacterClass = "Barbarian"
		hp := calculateHitPoints(testCharacter)
		assert.Equal(t, hp, 11)
	})

	t.Run("Calculate level", func(t *testing.T) {
		testCharacter.ExperiencePoints = 265000
		applyXpGains(&testCharacter, 1)
		assert.Equal(t, testCharacter.Level, 18)
	})
}
