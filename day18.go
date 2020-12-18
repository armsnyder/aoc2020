package main

import (
	"io"
	"strconv"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(18, func(part2 bool, inputReader io.Reader) interface{} {
	sum := 0
	aocutil.VisitStrings(inputReader, func(v string) {
		sum += day18Evaluate(v, part2)
	})
	return sum
})

func day18Evaluate(expression string, part2 bool) int {
	var operatorStack []rune
	var operandStack []int

	canReduceLastExpression := func(nextOperator rune) bool {
		if len(operatorStack) == 0 {
			return false
		}

		topOperator := operatorStack[len(operatorStack)-1]

		switch topOperator {
		case '(':
			return false
		case '*':
			return !part2
		case '+':
			return true
		default:
			panic(topOperator)
		}
	}

	reduceLastExpression := func() {
		topOperator := operatorStack[len(operatorStack)-1]
		topTwoOperands := operandStack[len(operandStack)-2:]

		var resultOperand int

		switch topOperator {
		case '+':
			resultOperand = topTwoOperands[0] + topTwoOperands[1]
		case '*':
			resultOperand = topTwoOperands[0] * topTwoOperands[1]
		default:
			panic(string(topOperator))
		}

		operandStack = append(operandStack[:len(operandStack)-2], resultOperand)
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	for _, ch := range expression {
		switch ch {
		case ' ':
		case '(':
			operatorStack = append(operatorStack, ch)
		case ')':
			for operatorStack[len(operatorStack)-1] != '(' {
				reduceLastExpression()
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		case '*', '+':
			if canReduceLastExpression(ch) {
				reduceLastExpression()
			}
			operatorStack = append(operatorStack, ch)
		default:
			operand, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}
			operandStack = append(operandStack, operand)
		}
	}

	for len(operandStack) > 1 {
		reduceLastExpression()
	}

	return operandStack[0]
}
