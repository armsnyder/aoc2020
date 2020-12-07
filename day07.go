package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(7, func(part2 bool, rawInput []byte) interface{} {
	myBag := "shiny gold"

	if part2 {
		tree := day07CreateBagTree(rawInput)
		memo := make(map[string]int)
		return day07SumNestedBagsInBag(myBag, tree, memo)
	}

	tree := day07CreateInverseBagTreeIgnoreQuantities(rawInput)
	return day07SumBagsThatCanContain(myBag, tree)
})

func day07CreateInverseBagTreeIgnoreQuantities(rawInput []byte) map[string][]string {
	tree := make(map[string][]string)
	bagRules := day07GetBagRules(rawInput)

	for _, bagRule := range bagRules {
		for _, bagQuantity := range bagRule.contents {
			tree[bagQuantity.bag] = append(tree[bagQuantity.bag], bagRule.containerBag)
		}
	}

	return tree
}

func day07SumBagsThatCanContain(bag string, tree map[string][]string) int {
	visited := make(map[string]bool)
	queue := []string{bag}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if visited[cur] {
			continue
		}

		visited[cur] = true

		for _, bag := range tree[cur] {
			queue = append(queue, bag)
		}
	}

	return len(visited) - 1
}

type day07BagQuantity struct {
	bag      string
	quantity int
}

func day07CreateBagTree(rawInput []byte) map[string][]day07BagQuantity {
	tree := make(map[string][]day07BagQuantity)
	bagRules := day07GetBagRules(rawInput)
	for _, bagRule := range bagRules {
		tree[bagRule.containerBag] = bagRule.contents
	}
	return tree
}

func day07SumNestedBagsInBag(bag string, tree map[string][]day07BagQuantity, memo map[string]int) int {
	if quantity, ok := memo[bag]; ok {
		return quantity
	}

	bagContents := tree[bag]
	total := 0

	for _, bagQuantity := range bagContents {
		containedBagsPerBag := day07SumNestedBagsInBag(bagQuantity.bag, tree, memo)
		totalContainedBagsIncludingItself := bagQuantity.quantity * (containedBagsPerBag + 1)

		total += totalContainedBagsIncludingItself
	}

	memo[bag] = total

	return total
}

type day07BagRule struct {
	containerBag string
	contents     []day07BagQuantity
}

func day07GetBagRules(rawInput []byte) []day07BagRule {
	contentsPattern := regexp.MustCompile(`(\d+) (\w+ \w+)`)
	var bagRules []day07BagRule
	rawBagRules := aocutil.Strings(rawInput)

	for _, rawBagRule := range rawBagRules {
		var bagRule day07BagRule
		bagRule.containerBag = strings.Join(strings.Fields(rawBagRule)[:2], " ")
		contentsSearch := strings.Split(rawBagRule, "contain")[1]
		contentsMatches := contentsPattern.FindAllStringSubmatch(contentsSearch, -1)

		for _, match := range contentsMatches {
			bag := match[2]
			quantity, _ := strconv.Atoi(match[1])

			bagRule.contents = append(bagRule.contents, day07BagQuantity{bag: bag, quantity: quantity})
		}

		bagRules = append(bagRules, bagRule)
	}

	return bagRules
}
