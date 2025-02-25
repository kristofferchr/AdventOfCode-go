// Code generated by aocgen; DO NOT EDIT.
package year2020

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay01(t *testing.T) {
	day01 := Day01{}
	input := []string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	}

	t.Run("partA", func(t *testing.T) {
		result := day01.PartA(input)

		assert.Equal(t, 514579, result)
	})

	t.Run("partB", func(t *testing.T) {
		result := day01.PartB(input)
		assert.Equal(t, 241861950, result)
	})
}
