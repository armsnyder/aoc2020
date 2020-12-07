package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/armsnyder/aoc2020/aocutil"
)

type dayFn func(part2 bool, rawInput []byte) interface{}

var days = make(map[int]dayFn)

func declareDay(day int, dayFn dayFn) interface{} {
	days[day] = dayFn
	return nil
}

func main() {
	easternTime, err := time.LoadLocation("EST")
	if err != nil {
		panic(err)
	}

	var (
		day        int
		part2      bool
		skipSubmit bool
		inputFile  string
	)
	flag.IntVar(&day, "d", time.Now().In(easternTime).Day(), "Day of the month")
	flag.BoolVar(&part2, "2", false, "Run part 2")
	flag.BoolVar(&skipSubmit, "s", false, "Skip submitting the answer")
	flag.StringVar(&inputFile, "f", "", "Puzzle input override filepath")
	flag.Parse()

	dayFn, ok := days[day]
	if !ok {
		fmt.Printf("Generating a stub for day %d\n", day)
		aocutil.GenerateStub(day)
		return
	}

	var input []byte
	if inputFile != "" {
		var err error
		input, err = ioutil.ReadFile(inputFile)
		if err != nil {
			panic(err)
		}
	} else {
		input = aocutil.GetInput(day)
	}

	start := time.Now()
	output := dayFn(part2, input)
	runTime := time.Since(start)

	fmt.Printf("Finished in %s\n", runTime)
	fmt.Println("Answer:", output)

	if !skipSubmit {
		aocutil.Submit(day, part2, output)
	}
}
