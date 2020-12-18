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
Benchmark/Day_01/Part_1-16  	  126088	      9191 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	  102786	     11826 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5661	    221147 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    5655	    210684 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   58038	     19327 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   41498	     28523 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    2240	    550039 ns/op	  271256 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1693	    695721 ns/op	  296961 B/op	    5872 allocs/op
Benchmark/Day_05/Part_1-16  	   16234	     73363 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    8000	    151007 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1424	    861593 ns/op	  330374 B/op	    4768 allocs/op
Benchmark/Day_06/Part_2-16  	    1383	    896432 ns/op	  330340 B/op	    4767 allocs/op
Benchmark/Day_07/Part_1-16  	     560	   1968012 ns/op	  776793 B/op	    8531 allocs/op
Benchmark/Day_07/Part_2-16  	     691	   1775848 ns/op	  689872 B/op	    7404 allocs/op
Benchmark/Day_08/Part_1-16  	   14199	     84359 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6315	    184834 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   15566	     77193 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3446	    356813 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  130501	      9175 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  119643	      9595 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      51	  21567798 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      40	  26151264 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   30848	     38070 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   33000	     36271 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  539529	      2049 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  444793	      2678 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    8938	    145833 ns/op	   62615 B/op	    1148 allocs/op
Benchmark/Day_14/Part_2-16  	     132	   8660513 ns/op	 5606080 B/op	    4329 allocs/op
Benchmark/Day_15/Part_1-16  	  222505	      4974 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       2	 556714657 ns/op	120008896 B/op	       7 allocs/op
Benchmark/Day_16/Part_1-16  	    6652	    175609 ns/op	  156616 B/op	     842 allocs/op
Benchmark/Day_16/Part_2-16  	    2685	    426067 ns/op	  174624 B/op	     981 allocs/op
Benchmark/Day_17/Part_1-16  	     178	   6779599 ns/op	  784422 B/op	     505 allocs/op
Benchmark/Day_17/Part_2-16  	       6	 196491408 ns/op	 9437913 B/op	   11153 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	50.370s
```
