package character

import (
	"reflect"
	"testing"
)

func TestCharacterValidations(t *testing.T) {

	t.Run("Valid Race", func(t *testing.T) {
		got := isValidRace("Human")
		checkResult(t, got, true)

	})

	t.Run("Valid Class", func(t *testing.T) {
		got := isValidClass("Bard")
		checkResult(t, got, true)

	})

	t.Run("Invalid Race", func(t *testing.T) {
		got := isValidRace("t-1000")
		checkResult(t, got, false)

	})

	t.Run("Invalid Class", func(t *testing.T) {
		got := isValidClass("terminator")
		checkResult(t, got, false)

	})

	//skipping name check for now
}

func checkResult(t testing.TB, got, want bool) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
