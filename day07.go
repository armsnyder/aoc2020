package main

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(7, func(part2 bool, inputReader io.Reader) interface{} {
	myBag := "shiny gold"

	if part2 {
		tree := day07CreateBagTree(inputReader)
		memo := make(map[string]int)
		return day07SumNestedBagsInBag(myBag, tree, memo)
	}

	tree := day07CreateInverseBagTreeIgnoreQuantities(inputReader)
	return day07SumBagsThatCanContain(myBag, tree)
})

func day07CreateInverseBagTreeIgnoreQuantities(inputReader io.Reader) map[string][]string {
	tree := make(map[string][]string)
	day07VisitBagRules(inputReader, func(containerBag string, contents []day07BagQuantity) {
		for _, bagQuantity := range contents {
			tree[bagQuantity.bag] = append(tree[bagQuantity.bag], containerBag)
		}
	})
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

func day07CreateBagTree(inputReader io.Reader) map[string][]day07BagQuantity {
	tree := make(map[string][]day07BagQuantity)
	day07VisitBagRules(inputReader, func(containerBag string, contents []day07BagQuantity) {
		tree[containerBag] = contents
	})
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

func day07VisitBagRules(inputReader io.Reader, visitFn func(containerBag string, contents []day07BagQuantity)) {
	contentsPattern := regexp.MustCompile(`(\d+) (\w+ \w+)`)

	aocutil.VisitStrings(inputReader, func(v string) {
		containerBag := strings.Join(strings.Fields(v)[:2], " ")
		contentsSearch := strings.Split(v, "contain")[1]
		contentsMatches := contentsPattern.FindAllStringSubmatch(contentsSearch, -1)

		var contents []day07BagQuantity

		for _, match := range contentsMatches {
			bag := match[2]
			quantity, _ := strconv.Atoi(match[1])

			contents = append(contents, day07BagQuantity{bag: bag, quantity: quantity})
		}

		visitFn(containerBag, contents)
	})
}
