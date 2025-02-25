// Code generated by aocgen; DO NOT EDIT.
package year2023

import (
	"image"
	"regexp"
	"strconv"
	"unicode"
)

type Day03 struct{}

type WholeNum struct {
	RowIndex           int
	FirstDigitColIndex int
	LastDigitColIndex  int
	Value              int
}

func (p Day03) PartA(lines []string) any {
	symbols := getAllSymbol(lines)
	parts := getAllParts(symbols, lines)

	sum := 0
	for _, part := range parts {
		for _, partNum := range part {
			sum += partNum
		}

	}
	return sum
}

func getAllSymbol(lines []string) map[image.Point]rune {
	schematic := map[image.Point]rune{}
	for rowIndex, line := range lines {
		if line == "" {
			continue
		}

		row := []rune{}
		for colIndex, r := range line {
			row = append(row, r)
			if r != '.' && !unicode.IsDigit(r) {
				schematic[image.Point{
					X: colIndex,
					Y: rowIndex,
				}] = r
			}
		}
	}

	return schematic

}

func getAllParts(symbols map[image.Point]rune, lines []string) map[image.Point][]int {
	parts := map[image.Point][]int{}
	for rowIndex, line := range lines {
		for _, m := range regexp.MustCompile("\\d+").FindAllStringIndex(line, -1) {
			bounds := map[image.Point]struct{}{}
			for colIndex := m[0]; colIndex < m[1]; colIndex++ {
				for _, d := range []image.Point{
					{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
				} {
					bounds[image.Point{colIndex, rowIndex}.Add(d)] = struct{}{}
				}
			}

			wholeNum, _ := strconv.Atoi(line[m[0]:m[1]])
			for pos := range bounds {
				if _, ok := symbols[pos]; ok {
					parts[pos] = append(parts[pos], wholeNum)
				}
			}
		}
	}
	return parts
}

func (p Day03) PartB(lines []string) any {
	symbols := getAllSymbol(lines)
	parts := getAllParts(symbols, lines)
	sum := 0
	for pos, part := range parts {
		if symbols[pos] == '*' && len(part) == 2 {
			sum += part[0] * part[1]
		}

	}
	return sum
}
