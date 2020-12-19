package main

import (
	"io"
	"strconv"
	"strings"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(19, func(part2 bool, inputReader io.Reader) interface{} {
	ruleSet, messages := day19Parse(inputReader)

	if part2 {
		ruleSet[8] = day19Rule{children: [][]int{{42}, {42, 8}}}
		ruleSet[11] = day19Rule{children: [][]int{{42, 31}, {42, 11, 31}}}
	}

	total := 0

	for _, message := range messages {
		matches := ruleSet.match(message, 0, 0, len(message))
		for _, match := range matches {
			if match == len(message) {
				total++
				break
			}
		}
	}

	return total
})

func day19Parse(inputReader io.Reader) (ruleSet day19RuleSet, messages []string) {
	ruleSet = make(day19RuleSet)

	aocutil.VisitStringGroups(inputReader, func(v []string) {
		if len(ruleSet) == 0 {
			for _, ruleRaw := range v {
				split1 := strings.SplitN(ruleRaw, ": ", 2)
				ruleID, err := strconv.Atoi(split1[0])
				if err != nil {
					panic(err)
				}

				if strings.HasPrefix(split1[1], `"`) {
					ruleSet[ruleID] = day19Rule{literal: strings.Trim(split1[1], `"`)}
					continue
				}

				split2 := strings.Split(split1[1], " | ")
				children := make([][]int, len(split2))

				for i, sub := range split2 {
					ids := strings.Fields(sub)
					child := make([]int, len(ids))

					for j, idStr := range ids {
						id, err := strconv.Atoi(idStr)
						if err != nil {
							panic(err)
						}
						child[j] = id
					}

					children[i] = child
				}
				ruleSet[ruleID] = day19Rule{children: children}
			}
		} else {
			messages = v
		}
	})

	return ruleSet, messages
}

type day19Rule struct {
	literal  string
	children [][]int
}

type day19RuleSet map[int]day19Rule

func (r day19RuleSet) match(s string, rid, depth, maxDepth int) (resultList []int) {
	if depth > maxDepth {
		return nil
	}

	if s == "" {
		return nil
	}

	rule := r[rid]

	if rule.children == nil {
		if strings.HasPrefix(s, rule.literal) {
			return []int{len(rule.literal)}
		}
		return nil
	}

	type job struct {
		partIndex  int
		matchSoFar int
	}

	resultSet := make(map[int]bool)

	depth++

	for _, child := range rule.children {
		seenJobs := make(map[job]bool)
		stack := []job{{}}

		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if seenJobs[cur] {
				continue
			}
			seenJobs[cur] = true

			if cur.partIndex == len(child) {
				resultSet[cur.matchSoFar] = true
				continue
			}

			matches := r.match(s[cur.matchSoFar:], child[cur.partIndex], depth, maxDepth)
			for _, match := range matches {
				stack = append(stack, job{
					partIndex:  cur.partIndex + 1,
					matchSoFar: cur.matchSoFar + match,
				})
			}
		}
	}

	if len(resultSet) == 0 {
		return nil
	}

	resultList = make([]int, 0, len(resultSet))
	for result := range resultSet {
		resultList = append(resultList, result)
	}

	return resultList
}
