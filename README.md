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
Benchmark/Day_01/Part_1-16  	  119476	      8661 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	  107115	     11380 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5370	    214655 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    6339	    204280 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   64208	     18648 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   41048	     29915 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    2251	    532149 ns/op	  271256 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1850	    634215 ns/op	  296961 B/op	    5872 allocs/op
Benchmark/Day_05/Part_1-16  	   17822	     68506 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    8187	    140318 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1464	    821504 ns/op	  330386 B/op	    4768 allocs/op
Benchmark/Day_06/Part_2-16  	    1446	    869197 ns/op	  330387 B/op	    4768 allocs/op
Benchmark/Day_07/Part_1-16  	     636	   1979108 ns/op	  776255 B/op	    8530 allocs/op
Benchmark/Day_07/Part_2-16  	     672	   1783972 ns/op	  690183 B/op	    7404 allocs/op
Benchmark/Day_08/Part_1-16  	   14864	     79471 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6922	    177392 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   16660	     72569 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3127	    364425 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  135438	      8686 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  124461	      8954 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      56	  21250738 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      49	  24629255 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   33861	     35551 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   35968	     33849 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  576949	      1966 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  477138	      2550 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    9232	    129009 ns/op	   62608 B/op	    1148 allocs/op
Benchmark/Day_14/Part_2-16  	     156	   7784337 ns/op	 5604830 B/op	    4324 allocs/op
Benchmark/Day_15/Part_1-16  	  235233	      4877 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       2	 507802212 ns/op	120008896 B/op	       7 allocs/op
Benchmark/Day_16/Part_1-16  	    7272	    166896 ns/op	  156616 B/op	     842 allocs/op
Benchmark/Day_16/Part_2-16  	    3037	    409437 ns/op	  174624 B/op	     981 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	46.128s
```
