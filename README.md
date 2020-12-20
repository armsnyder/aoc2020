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
Benchmark/Day_01/Part_1-16  	  116236	      9410 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	   89827	     12796 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5539	    219802 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    5924	    207054 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   59955	     19322 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   42322	     28761 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    2142	    555996 ns/op	  271256 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1838	    666837 ns/op	  296946 B/op	    5871 allocs/op
Benchmark/Day_05/Part_1-16  	   16197	     73539 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    7830	    149602 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1393	    870500 ns/op	  330311 B/op	    4767 allocs/op
Benchmark/Day_06/Part_2-16  	    1369	    885567 ns/op	  330393 B/op	    4768 allocs/op
Benchmark/Day_07/Part_1-16  	     612	   1959322 ns/op	  776486 B/op	    8531 allocs/op
Benchmark/Day_07/Part_2-16  	     670	   1827350 ns/op	  690419 B/op	    7405 allocs/op
Benchmark/Day_08/Part_1-16  	   14109	     84043 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6624	    184556 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   15432	     76112 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3430	    361043 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  129513	      9169 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  120489	      9676 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      55	  21792975 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      46	  25668289 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   32058	     38351 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   32865	     36315 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  565152	      2035 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  456051	      2641 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    9079	    135028 ns/op	   62612 B/op	    1148 allocs/op
Benchmark/Day_14/Part_2-16  	     140	   8484932 ns/op	 5604674 B/op	    4323 allocs/op
Benchmark/Day_15/Part_1-16  	  229710	      5011 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       2	 546240250 ns/op	120008896 B/op	       7 allocs/op
Benchmark/Day_16/Part_1-16  	    6898	    172032 ns/op	  156616 B/op	     842 allocs/op
Benchmark/Day_16/Part_2-16  	    2828	    436441 ns/op	  174624 B/op	     981 allocs/op
Benchmark/Day_17/Part_1-16  	     180	   6767281 ns/op	  784217 B/op	     502 allocs/op
Benchmark/Day_17/Part_2-16  	       6	 195516711 ns/op	 9451592 B/op	   11193 allocs/op
Benchmark/Day_18/Part_1-16  	    4924	    248697 ns/op	   77824 B/op	    6789 allocs/op
Benchmark/Day_18/Part_2-16  	    4413	    282014 ns/op	  105472 B/op	    7288 allocs/op
Benchmark/Day_19/Part_1-16  	     216	   5466527 ns/op	  761536 B/op	   83791 allocs/op
Benchmark/Day_19/Part_2-16  	      30	  39367472 ns/op	 5320964 B/op	  645596 allocs/op
Benchmark/Day_20/Part_1-16  	    2344	    508137 ns/op	  304612 B/op	    5077 allocs/op
Benchmark/Day_20/Part_2-16  	     630	   1922134 ns/op	  815368 B/op	   16666 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	57.157s
```
