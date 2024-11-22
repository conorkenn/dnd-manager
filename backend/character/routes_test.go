package character

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// invalid requests just testing if routing is hitting correct end point
func TestCharacterRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	CharacterRoutes(router)

	t.Run("Create Character Post /characters", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/characters", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Get Character Get /characters/:name", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/characters/name", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Get Characters Get /characters", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/characters", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	//skipping delete for now
}
