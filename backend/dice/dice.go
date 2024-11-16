package dice

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Dice struct {
	Sides    int `json:"sides"`
	NumRolls int `json:"num_rolls,omitempty"`
}

func Roll(c *gin.Context) {
	var newDice Dice
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	if err := c.BindJSON(&newDice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newDice.Sides <= 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The dice needs more than one side"})
		return
	}

	// optional field default to 1 if empty
	if newDice.NumRolls == 0 {
		newDice.NumRolls = 1
	}

	var results []int

	for i := 0; i < newDice.NumRolls; i++ {
		results = append(results, r.Intn(newDice.Sides)+1)
	}

	c.JSON(http.StatusOK, gin.H{
		"sides":     newDice.Sides,
		"num_rolls": newDice.NumRolls,
		"results":   results,
		"message":   "Successfully rolled dice",
	})
}
