package year2021

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay02(t *testing.T) {
	day02 := Day02{}
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	t.Run("partA", func(t *testing.T) {
		result := day02.PartA(input)

		assert.Equal(t, 150, result)
	})

	t.Run("partB", func(t *testing.T) {
		result := day02.PartB(input)
		assert.Equal(t, 900, result)
	})

}
