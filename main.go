package main

import (
	"flag"
	"fmt"
	"io"
	"time"

	"github.com/armsnyder/aoc2020/aocutil"
)

type dayFn func(part2 bool, inputReader io.Reader) interface{}

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
		day   int
		part2 bool
	)
	flag.IntVar(&day, "d", time.Now().In(easternTime).Day(), "Day of the month")
	flag.BoolVar(&part2, "2", false, "Run part 2")
	flag.Parse()

	dayFn, ok := days[day]
	if !ok {
		fmt.Printf("Generating a stub for day %d\n", day)
		aocutil.GenerateStub(day)
		return
	}

	input := aocutil.GetInput(day)
	defer input.Close()

	start := time.Now()
	output := dayFn(part2, input)
	runTime := time.Since(start)
	fmt.Printf("Finished in %s\n", runTime)

	aocutil.Submit(day, part2, output)
}
