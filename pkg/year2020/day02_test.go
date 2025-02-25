// Code generated by aocgen; DO NOT EDIT.
package year2020

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay02(t *testing.T) {
	day02 := Day02{}
	input := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	t.Run("partA", func(t *testing.T) {
		result := day02.PartA(input)

		assert.Equal(t, 2, result)
	})

	t.Run("partB", func(t *testing.T) {
		result := day02.PartB(input)
		assert.Equal(t, "implement_me", result)
	})
}
