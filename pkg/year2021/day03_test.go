package year2021

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay03(t *testing.T) {
	day03 := Day03{}
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	t.Run("partA", func(t *testing.T) {
		result := day03.PartA(input)

		assert.Equal(t, 198, result)
	})

	t.Run("partB", func(t *testing.T) {
		result := day03.PartB(input)
		assert.Equal(t, 230, result)
	})

}
