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
Benchmark/Day_01/Part_1-16  	  117799	      9736 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_01/Part_2-16  	   91126	     12762 ns/op	    9000 B/op	     211 allocs/op
Benchmark/Day_02/Part_1-16  	    5161	    229229 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_02/Part_2-16  	    5779	    224065 ns/op	   94760 B/op	    2002 allocs/op
Benchmark/Day_03/Part_1-16  	   52398	     23735 ns/op	   30800 B/op	     334 allocs/op
Benchmark/Day_03/Part_2-16  	   37614	     32103 ns/op	   30808 B/op	     335 allocs/op
Benchmark/Day_04/Part_1-16  	    1839	    612779 ns/op	  271257 B/op	    5340 allocs/op
Benchmark/Day_04/Part_2-16  	    1753	    675189 ns/op	  296943 B/op	    5871 allocs/op
Benchmark/Day_05/Part_1-16  	   16077	     75220 ns/op	   18248 B/op	     886 allocs/op
Benchmark/Day_05/Part_2-16  	    7863	    153418 ns/op	   34656 B/op	     898 allocs/op
Benchmark/Day_06/Part_1-16  	    1365	    903824 ns/op	  330350 B/op	    4768 allocs/op
Benchmark/Day_06/Part_2-16  	    1356	    902544 ns/op	  330360 B/op	    4768 allocs/op
Benchmark/Day_07/Part_1-16  	     590	   2020383 ns/op	  776943 B/op	    8531 allocs/op
Benchmark/Day_07/Part_2-16  	     642	   1886486 ns/op	  690297 B/op	    7405 allocs/op
Benchmark/Day_08/Part_1-16  	   13136	     89568 ns/op	   79040 B/op	    1272 allocs/op
Benchmark/Day_08/Part_2-16  	    6454	    186826 ns/op	  166720 B/op	    1409 allocs/op
Benchmark/Day_09/Part_1-16  	   15327	     78349 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_09/Part_2-16  	    3283	    364021 ns/op	   30688 B/op	    1004 allocs/op
Benchmark/Day_10/Part_1-16  	  119757	      9638 ns/op	    6472 B/op	     112 allocs/op
Benchmark/Day_10/Part_2-16  	  122560	     10236 ns/op	    7368 B/op	     113 allocs/op
Benchmark/Day_11/Part_1-16  	      52	  22135749 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_11/Part_2-16  	      45	  26086895 ns/op	   31928 B/op	     209 allocs/op
Benchmark/Day_12/Part_1-16  	   30484	     39307 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_12/Part_2-16  	   32404	     37524 ns/op	    6448 B/op	     792 allocs/op
Benchmark/Day_13/Part_1-16  	  544454	      2064 ns/op	    5464 B/op	       6 allocs/op
Benchmark/Day_13/Part_2-16  	  449709	      2690 ns/op	    5472 B/op	       7 allocs/op
Benchmark/Day_14/Part_1-16  	    9070	    138822 ns/op	   62614 B/op	    1148 allocs/op
Benchmark/Day_14/Part_2-16  	     133	   8975966 ns/op	 5605299 B/op	    4324 allocs/op
Benchmark/Day_15/Part_1-16  	  221593	      5263 ns/op	   12472 B/op	       7 allocs/op
Benchmark/Day_15/Part_2-16  	       2	 545681897 ns/op	120008896 B/op	       7 allocs/op
Benchmark/Day_16/Part_1-16  	    6831	    174944 ns/op	  156616 B/op	     842 allocs/op
Benchmark/Day_16/Part_2-16  	    3013	    439316 ns/op	  174624 B/op	     981 allocs/op
Benchmark/Day_17/Part_1-16  	     176	   7162837 ns/op	  783981 B/op	     502 allocs/op
Benchmark/Day_17/Part_2-16  	       5	 201651796 ns/op	 9437196 B/op	   11159 allocs/op
Benchmark/Day_18/Part_1-16  	    4767	    256177 ns/op	   77824 B/op	    6789 allocs/op
Benchmark/Day_18/Part_2-16  	    4160	    288859 ns/op	  105472 B/op	    7288 allocs/op
Benchmark/Day_19/Part_1-16  	     212	   5594045 ns/op	  761530 B/op	   83791 allocs/op
Benchmark/Day_19/Part_2-16  	      32	  40273973 ns/op	 5321244 B/op	  645597 allocs/op
Benchmark/Day_20/Part_1-16  	    2254	    520730 ns/op	  304614 B/op	    5077 allocs/op
Benchmark/Day_20/Part_2-16  	     603	   1957387 ns/op	  814491 B/op	   16643 allocs/op
Benchmark/Day_21/Part_1-16  	    2616	    441545 ns/op	  207195 B/op	     314 allocs/op
Benchmark/Day_21/Part_2-16  	    3874	    319470 ns/op	  192026 B/op	     317 allocs/op
Benchmark/Day_22/Part_1-16  	   88674	     13334 ns/op	    7088 B/op	     109 allocs/op
Benchmark/Day_22/Part_2-16  	       3	 431546153 ns/op	524906960 B/op	  495514 allocs/op
Benchmark/Day_23/Part_1-16  	  228391	      5358 ns/op	    4952 B/op	      18 allocs/op
Benchmark/Day_23/Part_2-16  	       1	2671827892 ns/op	117743264 B/op	 1038187 allocs/op
Benchmark/Day_24/Part_1-16  	    7231	    163882 ns/op	   47232 B/op	     510 allocs/op
Benchmark/Day_24/Part_2-16  	       7	 157721256 ns/op	23839570 B/op	   26148 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	72.152s
```
