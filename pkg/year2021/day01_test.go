package year2021

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay01(t *testing.T) {
	day01 := Day01{}
	sweepReport := []string{
		"199",
		"200",
		"208",
		"210",
		"200",
		"207",
		"240",
		"269",
		"260",
		"263",
	}
	t.Run("partA", func(t *testing.T) {

		result := day01.PartA(sweepReport)
		assert.Equal(t, 7, result)
	})
	t.Run("partB", func(t *testing.T) {
		result := day01.PartB(sweepReport)
		assert.Equal(t, 5, result)
	})

}
