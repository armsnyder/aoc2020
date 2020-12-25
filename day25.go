package main

import (
	"io"

	"github.com/armsnyder/aoc2020/aocutil"
)

var _ = declareDay(25, func(_ bool, inputReader io.Reader) interface{} {
	keys := aocutil.ReadAllInts(inputReader)
	key0LoopSize := day25CrackLoopSize(keys[0])
	return day25MakeEncryptionKey(keys[1], key0LoopSize)
})

func day25CrackLoopSize(publicKey int) int {
	loopSize := 0
	for value := 1; value != publicKey; loopSize++ {
		value = day25Transform(value, 7)
	}
	return loopSize
}

func day25MakeEncryptionKey(publicKey, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = day25Transform(value, publicKey)
	}
	return value
}

func day25Transform(value, subjectNumber int) int {
	value *= subjectNumber
	value %= 20201227
	return value
}
