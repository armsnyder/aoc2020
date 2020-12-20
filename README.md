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
Benchmark/Day_01/Part_1-16  	  124058	      9704 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	   81212	     12661 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5694	    221606 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    6004	    207643 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   62869	     19310 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   42411	     28099 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    2184	    545316 ns/op	  271256 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1830	    645138 ns/op	  296938 B/op	    5871 allocs/op
Benchmark/Day_05/Part_1-16  	   16586	     72465 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    8364	    147721 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1396	    858691 ns/op	  330362 B/op	    4768 allocs/op
Benchmark/Day_06/Part_2-16  	    1366	    881291 ns/op	  330373 B/op	    4768 allocs/op
Benchmark/Day_07/Part_1-16  	     634	   1924986 ns/op	  776611 B/op	    8531 allocs/op
Benchmark/Day_07/Part_2-16  	     698	   1739726 ns/op	  690253 B/op	    7404 allocs/op
Benchmark/Day_08/Part_1-16  	   14640	     83516 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6685	    183195 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   15957	     74691 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3280	    347406 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  128486	      9149 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  128259	      9272 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      56	  21432951 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      48	  25372240 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   30609	     38945 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   33165	     36012 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  578119	      2014 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  466838	      2587 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    9277	    132092 ns/op	   62619 B/op	    1149 allocs/op
Benchmark/Day_14/Part_2-16  	     150	   7959814 ns/op	 5603365 B/op	    4320 allocs/op
Benchmark/Day_15/Part_1-16  	  227328	      4916 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       2	 532258058 ns/op	120008896 B/op	       7 allocs/op
Benchmark/Day_16/Part_1-16  	    6682	    172370 ns/op	  156616 B/op	     842 allocs/op
Benchmark/Day_16/Part_2-16  	    2883	    424514 ns/op	  174624 B/op	     981 allocs/op
Benchmark/Day_17/Part_1-16  	     176	   6781320 ns/op	  783844 B/op	     501 allocs/op
Benchmark/Day_17/Part_2-16  	       6	 190255556 ns/op	 9435150 B/op	   11140 allocs/op
Benchmark/Day_18/Part_1-16  	    5025	    243244 ns/op	   77824 B/op	    6789 allocs/op
Benchmark/Day_18/Part_2-16  	    4569	    273510 ns/op	  105472 B/op	    7288 allocs/op
Benchmark/Day_19/Part_1-16  	     219	   5421554 ns/op	  761509 B/op	   83791 allocs/op
Benchmark/Day_19/Part_2-16  	      28	  38719399 ns/op	 5321008 B/op	  645596 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	54.935s
```
