// Code generated by aocgen; DO NOT EDIT.
package year2023

import (
	"math"
	"regexp"
	"strings"
)

type Day04 struct{}

type ScratchCard struct {
	WinningNumbers map[string]struct{}
	MyNumbers []string
}

func (p Day04) PartA(lines []string) any {
	cards := parseScratchCards(lines)
	sum := 0
	for _, card := range cards {
		amountOfWinningNums := 0

		for _, myNum := range card.MyNumbers {
			if _, ok := card.WinningNumbers[myNum]; ok {
				amountOfWinningNums++
			}
		}
		sum += int(math.Pow(2, float64(amountOfWinningNums-1)))
	}

	return sum
}


func (p Day04) PartB(lines []string) any {
	cards := parseScratchCards(lines)
	copies := []ScratchCard{}

	for i, card := range cards {
		copies = winCopies(card, copies, cards, i)
	}

	return len(copies) + len(cards)
}

func parseScratchCards(lines []string) []ScratchCard {
	cards := []ScratchCard{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		removedCardNum := strings.TrimSpace(regexp.MustCompile("^\\w+ \\d+: ").ReplaceAllString(line, ""))
		winningNumsAndMyNums := strings.Split(removedCardNum, "|")
		winningNums := strings.TrimSpace(winningNumsAndMyNums[0])
		myNums := strings.TrimSpace(winningNumsAndMyNums[1])

		winningNumsMap := map[string]struct{}{}
		for _, winninNum := range strings.Fields(winningNums) {
			winningNumsMap[winninNum] = struct{}{}
		}
		cards = append(cards, ScratchCard{
			WinningNumbers: winningNumsMap,
			MyNumbers:      strings.Fields(myNums),
		})

	}
	return cards
}

func winCopies(currentCard ScratchCard, copies []ScratchCard, originalCards []ScratchCard,cardNum int) []ScratchCard{
	amountOfWinningNums := 0
	for _, myNum := range currentCard.MyNumbers {
		if _, ok := currentCard.WinningNumbers[myNum]; ok {
			amountOfWinningNums++
		}
	}

	for i := cardNum + 1;i <  cardNum  + 1 + amountOfWinningNums; i++  {
		copiedCard := originalCards[i]
		copies = append(copies, copiedCard)
		copies = winCopies(copiedCard, copies, originalCards, i)
	}

	return copies
}
