package character

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetCharacter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	characters = []Character{
		{Name: "Suel", Race: "Human", CharacterClass: "Wizard", Level: 31},
	}

	router.GET("/characters/:name", GetCharacter)

	t.Run("Found Character", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/characters/Suel", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		expectedResponse := map[string]interface{}{
			"name":           "Suel",
			"race":           "Human",
			"characterClass": "Wizard",
			"level":          31,
			"attributes": map[string]interface{}{
				"strength":     0,
				"dexterity":    0,
				"constitution": 0,
				"intelligence": 0,
				"wisdom":       0,
				"charisma":     0,
			},
		}

		expectedJSON, err := json.Marshal(expectedResponse)
		if err != nil {
			t.Fatalf("Failed to marshal expected response: %v", err)
		}

		assert.JSONEq(t, string(expectedJSON), w.Body.String())
	})

	t.Run("Character Not Found", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/characters/Suelchi", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)

		expectedJSON := `{"message":"character not found"}`

		assert.JSONEq(t, expectedJSON, w.Body.String())
	})
	t.Run("Failed to Load Characters", func(t *testing.T) {
		originalLoadFunc := loadCharactersFromCSV

		defer func() {
			loadCharactersFromCSV = originalLoadFunc
		}()

		loadCharactersFromCSV = func() error {
			return fmt.Errorf("failed to load characters")
		}

		req, _ := http.NewRequest(http.MethodGet, "/characters/Suelchi", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		expectedJSON := `{"error":"failed to load characters"}`
		assert.JSONEq(t, expectedJSON, w.Body.String())
	})

}
