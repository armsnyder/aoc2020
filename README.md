# aoc2020

My [Advent of Code 2020](https://adventofcode.com/2020) solutions.

Usage:

```shell
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
Benchmark/Day_01/Part_1-16  	  119534	      9236 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	  101847	     11525 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5508	    224858 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    5751	    213924 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   61533	     20474 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   40292	     29375 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    2168	    559111 ns/op	  271256 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1712	    672812 ns/op	  296944 B/op	    5871 allocs/op
Benchmark/Day_05/Part_1-16  	   16333	     73975 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    8386	    153180 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1396	    862463 ns/op	  330426 B/op	    4769 allocs/op
Benchmark/Day_06/Part_2-16  	    1410	    882457 ns/op	  330453 B/op	    4770 allocs/op
Benchmark/Day_07/Part_1-16  	     571	   2015157 ns/op	  776086 B/op	    8530 allocs/op
Benchmark/Day_07/Part_2-16  	     656	   1836020 ns/op	  691157 B/op	    7405 allocs/op
Benchmark/Day_08/Part_1-16  	   13753	     82480 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6619	    188086 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   15736	     75196 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3165	    379661 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  130377	      9083 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  125991	      9196 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      54	  21930367 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      45	  25850478 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   31195	     37552 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   34176	     35671 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  563704	      1989 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  461443	      2607 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    9158	    135289 ns/op	   62620 B/op	    1148 allocs/op
Benchmark/Day_14/Part_2-16  	     151	   7994621 ns/op	 5606474 B/op	    4331 allocs/op
Benchmark/Day_15/Part_1-16  	  216478	      5029 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       3	 428744444 ns/op	120008928 B/op	       7 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	43.950s
```
