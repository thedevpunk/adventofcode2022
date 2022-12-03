package day03

import (
	"fmt"
	"strings"
)

func Main() {
	var rucksacks = strings.Split(Input, "\n")
	part01(rucksacks)
	part02(rucksacks)
}

func part02(rucksacks []string) {
	total := 0
	groups := len(rucksacks)

	for i := 0; i < groups; i += 3 {
		r1 := rucksacks[i]
		r2 := rucksacks[i+1]
		r3 := rucksacks[i+2]

		groupBatch := "0"

		for _, letter := range r1 {
			if strings.Contains(r2, string(letter)) && strings.Contains(r3, string(letter)) {
				groupBatch = string(letter)
				break
			}
		}

		total += strings.Index(Values, groupBatch)
	}

	fmt.Printf("PART 02 - Total: %d\n", total)
}

func part01(rucksacks []string) {
	total := 0

	for _, rucksack := range rucksacks {
		length := len(rucksack)
		half := length / 2

		part1 := rucksack[0:half]
		part2 := rucksack[half:length]

		letterInBoth := isInBoth(part1, part2)

		points := strings.Index(Values, letterInBoth)

		total += points
	}

	fmt.Printf("PART 01 - Total: %d\n", total)
}

func isInBoth(left, right string) (letter string) {
	for _, letterLeft := range left {
		if strings.Contains(right, string(letterLeft)) {
			return string(letterLeft)
		}
	}
	return "0"
}
