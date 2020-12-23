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
Benchmark/Day_01/Part_1-16  	  119650	      9141 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	   97771	     12173 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5751	    213587 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    6105	    207885 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   62982	     19192 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   42184	     27664 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    2076	    559395 ns/op	  271256 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1722	    652359 ns/op	  296936 B/op	    5871 allocs/op
Benchmark/Day_05/Part_1-16  	   16934	     70675 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    8650	    146417 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1402	    856037 ns/op	  330400 B/op	    4769 allocs/op
Benchmark/Day_06/Part_2-16  	    1399	    859880 ns/op	  330381 B/op	    4768 allocs/op
Benchmark/Day_07/Part_1-16  	     595	   1896181 ns/op	  776040 B/op	    8530 allocs/op
Benchmark/Day_07/Part_2-16  	     670	   1756870 ns/op	  690456 B/op	    7404 allocs/op
Benchmark/Day_08/Part_1-16  	   14143	     82139 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6781	    180615 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   16059	     74803 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3598	    352652 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  132955	      8981 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  125114	      9392 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      56	  21277784 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      44	  25685439 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   33198	     36652 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   33256	     35393 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  589125	      2005 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  467756	      2619 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    9286	    131497 ns/op	   62619 B/op	    1149 allocs/op
Benchmark/Day_14/Part_2-16  	     150	   7948045 ns/op	 5606285 B/op	    4332 allocs/op
Benchmark/Day_15/Part_1-16  	  234202	      4808 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       3	 443285567 ns/op	120008896 B/op	       7 allocs/op
Benchmark/Day_16/Part_1-16  	    7075	    173797 ns/op	  156616 B/op	     842 allocs/op
Benchmark/Day_16/Part_2-16  	    3084	    408503 ns/op	  174624 B/op	     981 allocs/op
Benchmark/Day_17/Part_1-16  	     175	   6561118 ns/op	  784788 B/op	     506 allocs/op
Benchmark/Day_17/Part_2-16  	       6	 195909538 ns/op	 9445066 B/op	   11189 allocs/op
Benchmark/Day_18/Part_1-16  	    5110	    241473 ns/op	   77824 B/op	    6789 allocs/op
Benchmark/Day_18/Part_2-16  	    4575	    282759 ns/op	  105472 B/op	    7288 allocs/op
Benchmark/Day_19/Part_1-16  	     224	   5309832 ns/op	  761501 B/op	   83791 allocs/op
Benchmark/Day_19/Part_2-16  	      31	  38669511 ns/op	 5321167 B/op	  645597 allocs/op
Benchmark/Day_20/Part_1-16  	    2389	    488837 ns/op	  304621 B/op	    5077 allocs/op
Benchmark/Day_20/Part_2-16  	     630	   1933560 ns/op	  814324 B/op	   16641 allocs/op
Benchmark/Day_21/Part_1-16  	    2840	    414167 ns/op	  207189 B/op	     314 allocs/op
Benchmark/Day_21/Part_2-16  	    4171	    298016 ns/op	  192023 B/op	     317 allocs/op
Benchmark/Day_22/Part_1-16  	   99841	     11914 ns/op	    7088 B/op	     109 allocs/op
Benchmark/Day_22/Part_2-16  	       3	 415664122 ns/op	525201581 B/op	  496289 allocs/op
Benchmark/Day_23/Part_1-16  	  227337	      5123 ns/op	    4952 B/op	      18 allocs/op
Benchmark/Day_23/Part_2-16  	       1	2564895811 ns/op	117795400 B/op	 1038550 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	67.854s
```
