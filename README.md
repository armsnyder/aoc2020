# aoc2020

My [Advent of Code 2020](https://adventofcode.com/2020) solutions.

Usage:

```
$ go run . [-d day] [-2]
```

Requires a `session.txt` file containing a session token, for pulling puzzle input and submitting answers.
(Inputs and answers are cached.)

## Benchmarks

I thought it would be fun to share performance [benchmarks](https://golang.org/pkg/testing/#hdr-Benchmarks)
for each of my puzzle solutions, since I write benchmarks anyway to help guide my optimizations.
I don't always optimize for the best possible time if I think it impacts code readability.
Benchmarks use the real puzzle input, which is preloaded in memory.

```
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/armsnyder/aoc2020
Benchmark/Day_01/Part_1-16  	  117361	      9276 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	   96600	     11905 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5041	    230160 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    5870	    208583 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   61236	     19577 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   40929	     28426 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    2121	    553131 ns/op	  271257 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1828	    668776 ns/op	  296967 B/op	    5872 allocs/op
Benchmark/Day_05/Part_1-16  	   16514	     73392 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    7377	    151117 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1371	    884245 ns/op	  330400 B/op	    4769 allocs/op
Benchmark/Day_06/Part_2-16  	    1275	    893128 ns/op	  330336 B/op	    4767 allocs/op
Benchmark/Day_07/Part_1-16  	     609	   1967506 ns/op	  776718 B/op	    8530 allocs/op
Benchmark/Day_07/Part_2-16  	     644	   1810747 ns/op	  691055 B/op	    7405 allocs/op
Benchmark/Day_08/Part_1-16  	   14172	     84550 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6585	    188348 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   15582	     76247 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3429	    362430 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  124532	      9196 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  123232	      9475 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      55	  21612245 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      46	  25566112 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   30366	     39469 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   32629	     37364 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  599839	      2024 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  457795	      2634 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    9038	    134899 ns/op	   62611 B/op	    1148 allocs/op
Benchmark/Day_14/Part_2-16  	     140	   8212103 ns/op	 5606537 B/op	    4327 allocs/op
Benchmark/Day_15/Part_1-16  	  237426	      4984 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       2	 533661620 ns/op	120008896 B/op	       7 allocs/op
Benchmark/Day_16/Part_1-16  	    6924	    173731 ns/op	  156616 B/op	     842 allocs/op
Benchmark/Day_16/Part_2-16  	    2800	    420346 ns/op	  174624 B/op	     981 allocs/op
Benchmark/Day_17/Part_1-16  	     178	   6645471 ns/op	  785273 B/op	     509 allocs/op
Benchmark/Day_17/Part_2-16  	       6	 196202242 ns/op	 9434417 B/op	   11153 allocs/op
Benchmark/Day_18/Part_1-16  	    4808	    246708 ns/op	   77824 B/op	    6789 allocs/op
Benchmark/Day_18/Part_2-16  	    4435	    278913 ns/op	  105472 B/op	    7288 allocs/op
Benchmark/Day_19/Part_1-16  	     214	   5495415 ns/op	  761520 B/op	   83791 allocs/op
Benchmark/Day_19/Part_2-16  	      30	  39618680 ns/op	 5321187 B/op	  645597 allocs/op
Benchmark/Day_20/Part_1-16  	    2320	    510323 ns/op	  304610 B/op	    5077 allocs/op
Benchmark/Day_20/Part_2-16  	     621	   1941616 ns/op	  814527 B/op	   16644 allocs/op
Benchmark/Day_21/Part_1-16  	    2828	    437397 ns/op	  207189 B/op	     314 allocs/op
Benchmark/Day_21/Part_2-16  	    4117	    303084 ns/op	  192024 B/op	     317 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	60.345s
```
