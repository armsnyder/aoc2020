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
Benchmark/Day_01/Part_1-16  	  129750	      8744 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	  101368	     11707 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5528	    217603 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    5440	    200754 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   64213	     18718 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   43182	     27596 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    2239	    524482 ns/op	  271256 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1893	    636939 ns/op	  296917 B/op	    5871 allocs/op
Benchmark/Day_05/Part_1-16  	   17589	     70321 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    8407	    143071 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1444	    837931 ns/op	  330393 B/op	    4768 allocs/op
Benchmark/Day_06/Part_2-16  	    1423	    861294 ns/op	  330299 B/op	    4767 allocs/op
Benchmark/Day_07/Part_1-16  	     630	   1870731 ns/op	  777208 B/op	    8531 allocs/op
Benchmark/Day_07/Part_2-16  	     699	   1745754 ns/op	  690289 B/op	    7404 allocs/op
Benchmark/Day_08/Part_1-16  	   14882	     79603 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6789	    178041 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   15958	     74100 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3429	    345194 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  128862	      8758 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  133730	      8963 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      57	  21475349 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      49	  25000731 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   32919	     36039 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   33618	     35349 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  589862	      1979 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  474645	      2556 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    9146	    128273 ns/op	   62616 B/op	    1148 allocs/op
Benchmark/Day_14/Part_2-16  	     153	   7856588 ns/op	 5605913 B/op	    4329 allocs/op
Benchmark/Day_15/Part_1-16  	  232204	      4927 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       3	 445374721 ns/op	120008928 B/op	       7 allocs/op
Benchmark/Day_16/Part_1-16  	    6850	    171650 ns/op	  156616 B/op	     842 allocs/op
Benchmark/Day_16/Part_2-16  	    2973	    427891 ns/op	  174624 B/op	     981 allocs/op
Benchmark/Day_17/Part_1-16  	     184	   6706264 ns/op	  784062 B/op	     500 allocs/op
Benchmark/Day_17/Part_2-16  	       6	 193809805 ns/op	 9433286 B/op	   11140 allocs/op
Benchmark/Day_18/Part_1-16  	    5010	    240842 ns/op	   77824 B/op	    6789 allocs/op
Benchmark/Day_18/Part_2-16  	    4618	    274341 ns/op	  105472 B/op	    7288 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	52.246s
```
