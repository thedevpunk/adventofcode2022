package day02

import (
	"fmt"
	"strings"
)

func Main() {
	var rucksacks = strings.Split(Input, "\n")

	total := 0

	for _, rucksack := range rucksacks {

		length := len(rucksack)
		middle := length / 2 - 1

		part1 := rucksack[0:middle]
		part2 := rucksack[middle:length]

		letterInBoth := isInBoth(part1, part2)

		total += strings.Index(Values, letterInBoth)
	}

	fmt.Println("Total")
	fmt.Println(total)
}

func isInBoth(left, right string) (letter string) {
	fmt.Println("----");
	fmt.Println(left);
	fmt.Println(right);

	for _, letterLeft := range left {
		for _, letterRight := range right {
			if letterLeft == letterRight {
				fmt.Println(letterLeft)
				return string(letterLeft)
			}
		}
	}
	return "0"
}