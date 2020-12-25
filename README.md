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
Benchmark/Day_01/Part_1-16  	  130648	      8916 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	  103280	     11475 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5865	    217267 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    6362	    199416 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   63758	     18783 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   43674	     27196 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    2265	    525093 ns/op	  271257 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1746	    627896 ns/op	  296963 B/op	    5872 allocs/op
Benchmark/Day_05/Part_1-16  	   17012	     69634 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    8811	    143582 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1435	    832672 ns/op	  330315 B/op	    4767 allocs/op
Benchmark/Day_06/Part_2-16  	    1416	    843136 ns/op	  330354 B/op	    4768 allocs/op
Benchmark/Day_07/Part_1-16  	     640	   1893393 ns/op	  775977 B/op	    8530 allocs/op
Benchmark/Day_07/Part_2-16  	     718	   1720302 ns/op	  690236 B/op	    7404 allocs/op
Benchmark/Day_08/Part_1-16  	   14809	     79866 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6999	    177714 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   16597	     72392 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3644	    336142 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  131686	      8653 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  133900	      8964 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      57	  20705225 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      50	  24764981 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   33690	     36171 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   33578	     34512 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  606169	      1960 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  480672	      2529 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    9235	    127521 ns/op	   62611 B/op	    1148 allocs/op
Benchmark/Day_14/Part_2-16  	     162	   7333065 ns/op	 5606823 B/op	    4329 allocs/op
Benchmark/Day_15/Part_1-16  	  212142	      4811 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       3	 397788055 ns/op	120008896 B/op	       7 allocs/op
Benchmark/Day_16/Part_1-16  	    7147	    165010 ns/op	  156616 B/op	     842 allocs/op
Benchmark/Day_16/Part_2-16  	    3004	    417441 ns/op	  174624 B/op	     981 allocs/op
Benchmark/Day_17/Part_1-16  	     184	   6336696 ns/op	  783494 B/op	     500 allocs/op
Benchmark/Day_17/Part_2-16  	       6	 187683238 ns/op	 9431910 B/op	   11139 allocs/op
Benchmark/Day_18/Part_1-16  	    4677	    235040 ns/op	   77824 B/op	    6789 allocs/op
Benchmark/Day_18/Part_2-16  	    4588	    266378 ns/op	  105472 B/op	    7288 allocs/op
Benchmark/Day_19/Part_1-16  	     231	   5345259 ns/op	  761510 B/op	   83791 allocs/op
Benchmark/Day_19/Part_2-16  	      30	  38883044 ns/op	 5320919 B/op	  645596 allocs/op
Benchmark/Day_20/Part_1-16  	    2532	    482169 ns/op	  304594 B/op	    5077 allocs/op
Benchmark/Day_20/Part_2-16  	     666	   1845760 ns/op	  816217 B/op	   16694 allocs/op
Benchmark/Day_21/Part_1-16  	    2940	    417818 ns/op	  207199 B/op	     314 allocs/op
Benchmark/Day_21/Part_2-16  	    4034	    291154 ns/op	  192021 B/op	     317 allocs/op
Benchmark/Day_22/Part_1-16  	   92588	     11710 ns/op	    7088 B/op	     109 allocs/op
Benchmark/Day_22/Part_2-16  	       3	 412108979 ns/op	525103738 B/op	  496231 allocs/op
Benchmark/Day_23/Part_1-16  	  243758	      4982 ns/op	    4952 B/op	      18 allocs/op
Benchmark/Day_23/Part_2-16  	       1	2496229776 ns/op	117737776 B/op	 1038140 allocs/op
Benchmark/Day_24/Part_1-16  	    8930	    131896 ns/op	   47231 B/op	     510 allocs/op
Benchmark/Day_24/Part_2-16  	       7	 164353516 ns/op	23816508 B/op	   25916 allocs/op
Benchmark/Day_25/Part_1-16  	      31	  45685982 ns/op	    4144 B/op	       6 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	72.118s
```
