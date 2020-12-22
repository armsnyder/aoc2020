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
Benchmark/Day_01/Part_1-16  	  122793	      9047 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	   92320	     12195 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5731	    217708 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    6313	    200056 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   62503	     18984 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   43080	     28056 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    2202	    545004 ns/op	  271256 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1880	    637810 ns/op	  296956 B/op	    5872 allocs/op
Benchmark/Day_05/Part_1-16  	   16923	     71216 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    8413	    146538 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1388	    844290 ns/op	  330317 B/op	    4767 allocs/op
Benchmark/Day_06/Part_2-16  	    1362	    861063 ns/op	  330298 B/op	    4767 allocs/op
Benchmark/Day_07/Part_1-16  	     606	   1924802 ns/op	  775714 B/op	    8530 allocs/op
Benchmark/Day_07/Part_2-16  	     672	   1810453 ns/op	  690532 B/op	    7405 allocs/op
Benchmark/Day_08/Part_1-16  	   14324	     83607 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6841	    181082 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   16201	     73841 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3567	    348740 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  132246	      9140 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  127454	      9311 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      56	  21545928 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      46	  24919612 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   32796	     36677 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   32882	     35755 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  534904	      1998 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  467632	      2559 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    9256	    129823 ns/op	   62618 B/op	    1148 allocs/op
Benchmark/Day_14/Part_2-16  	     150	   7636697 ns/op	 5606299 B/op	    4329 allocs/op
Benchmark/Day_15/Part_1-16  	  237626	      4842 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       3	 440319574 ns/op	120008896 B/op	       7 allocs/op
Benchmark/Day_16/Part_1-16  	    7002	    167036 ns/op	  156616 B/op	     842 allocs/op
Benchmark/Day_16/Part_2-16  	    3129	    415575 ns/op	  174624 B/op	     981 allocs/op
Benchmark/Day_17/Part_1-16  	     184	   6543368 ns/op	  784134 B/op	     504 allocs/op
Benchmark/Day_17/Part_2-16  	       6	 191306715 ns/op	 9430645 B/op	   11126 allocs/op
Benchmark/Day_18/Part_1-16  	    5059	    240845 ns/op	   77824 B/op	    6789 allocs/op
Benchmark/Day_18/Part_2-16  	    4538	    270596 ns/op	  105472 B/op	    7288 allocs/op
Benchmark/Day_19/Part_1-16  	     223	   5317565 ns/op	  761602 B/op	   83792 allocs/op
Benchmark/Day_19/Part_2-16  	      31	  39015393 ns/op	 5321065 B/op	  645597 allocs/op
Benchmark/Day_20/Part_1-16  	    2449	    496368 ns/op	  304584 B/op	    5077 allocs/op
Benchmark/Day_20/Part_2-16  	     646	   1897875 ns/op	  813659 B/op	   16622 allocs/op
Benchmark/Day_21/Part_1-16  	    2926	    416607 ns/op	  207188 B/op	     314 allocs/op
Benchmark/Day_21/Part_2-16  	    4190	    298466 ns/op	  192031 B/op	     317 allocs/op
Benchmark/Day_22/Part_1-16  	  101455	     11925 ns/op	    7088 B/op	     109 allocs/op
Benchmark/Day_22/Part_2-16  	       3	 414598319 ns/op	525037370 B/op	  495937 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	64.351s
```
