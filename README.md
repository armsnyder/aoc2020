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
Benchmarks use the real puzzle input and include file I/O with the puzzle input file.
The `*BaselineIO` benchmarks read and discard the input file without running the solution code,
to serve as a baseline for the solution benchmarks.

```
$ go test -run=XXX -bench=.
goos: darwin
goarch: amd64
pkg: github.com/armsnyder/aoc2020
BenchmarkDay01BaselineIO-16    	   47641	     25778 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay01Part1-16         	   31830	     36521 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay01Part2-16         	   29839	     39669 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay02BaselineIO-16    	   41438	     27767 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay02Part1-16         	    4546	    268168 ns/op	   95181 B/op	    2009 allocs/op
BenchmarkDay02Part2-16         	    4920	    251153 ns/op	   95181 B/op	    2009 allocs/op
BenchmarkDay03BaselineIO-16    	   44871	     26409 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay03Part1-16         	   23775	     50902 ns/op	   31186 B/op	     341 allocs/op
BenchmarkDay03Part2-16         	   19687	     59930 ns/op	   31194 B/op	     342 allocs/op
BenchmarkDay04BaselineIO-16    	   42463	     28442 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay04Part1-16         	    1808	    615812 ns/op	  271772 B/op	    5347 allocs/op
BenchmarkDay04Part2-16         	    1692	    717430 ns/op	  297494 B/op	    5879 allocs/op
BenchmarkDay05BaselineIO-16    	   45174	     26701 ns/op	     371 B/op	       7 allocs/op
BenchmarkDay05Part1-16         	   10000	    105377 ns/op	   18626 B/op	     893 allocs/op
BenchmarkDay05Part2-16         	    6648	    189589 ns/op	   35036 B/op	     905 allocs/op
BenchmarkDay06BaselineIO-16    	   42732	     27871 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay06Part1-16         	    1303	    923467 ns/op	  330813 B/op	    4773 allocs/op
BenchmarkDay06Part2-16         	    1266	    948239 ns/op	  330924 B/op	    4775 allocs/op
BenchmarkDay07BaselineIO-16    	   38151	     32650 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay07Part1-16         	     576	   2014918 ns/op	  778454 B/op	    8538 allocs/op
BenchmarkDay07Part2-16         	     655	   1831459 ns/op	  691202 B/op	    7412 allocs/op
BenchmarkDay08BaselineIO-16    	   46746	     25756 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay08Part1-16         	    9860	    119025 ns/op	   79452 B/op	    1279 allocs/op
BenchmarkDay08Part2-16         	    5535	    218253 ns/op	  167185 B/op	    1416 allocs/op
BenchmarkDay09BaselineIO-16    	   44743	     26521 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay09Part1-16         	   10000	    111081 ns/op	   31073 B/op	    1011 allocs/op
BenchmarkDay09Part2-16         	    2911	    428299 ns/op	   31074 B/op	    1011 allocs/op
BenchmarkDay10BaselineIO-16    	   49221	     25491 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay10Part1-16         	   31869	     37854 ns/op	    6844 B/op	     119 allocs/op
BenchmarkDay10Part2-16         	   31650	     37153 ns/op	    7740 B/op	     120 allocs/op
BenchmarkDay11BaselineIO-16    	   45237	     26860 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay11Part1-16         	      34	  34428937 ns/op	 5442320 B/op	   58535 allocs/op
BenchmarkDay11Part2-16         	      24	  43494901 ns/op	10083034 B/op	   81990 allocs/op
BenchmarkDay12BaselineIO-16    	   47756	     25362 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay12Part1-16         	   17784	     66506 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay12Part2-16         	   18282	     64771 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay13BaselineIO-16    	   49021	     24976 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay13Part1-16         	   40813	     29662 ns/op	    5835 B/op	      13 allocs/op
BenchmarkDay13Part2-16         	   39504	     30150 ns/op	    5843 B/op	      14 allocs/op
BenchmarkDay14BaselineIO-16    	   42664	     26969 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay14Part1-16         	    7071	    177099 ns/op	   63020 B/op	    1156 allocs/op
BenchmarkDay14Part2-16         	     135	   8357183 ns/op	 5609123 B/op	    4340 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	60.511s
```
