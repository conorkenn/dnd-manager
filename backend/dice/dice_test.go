package dice

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRoll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/roll", Roll)

	t.Run("Roll Test", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, "/roll", strings.NewReader(`{"sides": 6, "num_rolls": 1}`))
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "message")
	})
}
