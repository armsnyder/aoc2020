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
BenchmarkDay01BaselineIO-16    	   47350	     25597 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay01Part1-16         	   29114	     38685 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay01Part2-16         	   29893	     40515 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay02BaselineIO-16    	   42818	     28166 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay02Part1-16         	    4796	    263990 ns/op	   95181 B/op	    2009 allocs/op
BenchmarkDay02Part2-16         	    5103	    245047 ns/op	   95181 B/op	    2009 allocs/op
BenchmarkDay03BaselineIO-16    	   44403	     26556 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay03Part1-16         	   23962	     51362 ns/op	   31186 B/op	     341 allocs/op
BenchmarkDay03Part2-16         	   19994	     59624 ns/op	   31194 B/op	     342 allocs/op
BenchmarkDay04BaselineIO-16    	   40240	     27682 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay04Part1-16         	    1772	    607042 ns/op	  271771 B/op	    5347 allocs/op
BenchmarkDay04Part2-16         	    1676	    723818 ns/op	  297475 B/op	    5879 allocs/op
BenchmarkDay05BaselineIO-16    	   45694	     26452 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay05Part1-16         	   10000	    105824 ns/op	   18626 B/op	     893 allocs/op
BenchmarkDay05Part2-16         	    6775	    186305 ns/op	   35035 B/op	     905 allocs/op
BenchmarkDay06BaselineIO-16    	   41415	     28663 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay06Part1-16         	    1281	    929768 ns/op	  330870 B/op	    4774 allocs/op
BenchmarkDay06Part2-16         	    1212	    943138 ns/op	  330878 B/op	    4775 allocs/op
BenchmarkDay07BaselineIO-16    	   36037	     31506 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay07Part1-16         	     578	   1972629 ns/op	  778657 B/op	    8538 allocs/op
BenchmarkDay07Part2-16         	     632	   1836093 ns/op	  690880 B/op	    7412 allocs/op
BenchmarkDay08BaselineIO-16    	   47176	     25327 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay08Part1-16         	    9340	    119522 ns/op	   79453 B/op	    1279 allocs/op
BenchmarkDay08Part2-16         	    5482	    217915 ns/op	  167186 B/op	    1416 allocs/op
BenchmarkDay09BaselineIO-16    	   44919	     26470 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay09Part1-16         	   10000	    109333 ns/op	   31073 B/op	    1011 allocs/op
BenchmarkDay09Part2-16         	    2904	    423538 ns/op	   31074 B/op	    1011 allocs/op
BenchmarkDay10BaselineIO-16    	   48002	     25015 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay10Part1-16         	   32738	     37660 ns/op	    6844 B/op	     119 allocs/op
BenchmarkDay10Part2-16         	   32331	     37779 ns/op	    7740 B/op	     120 allocs/op
BenchmarkDay11BaselineIO-16    	   46198	     26822 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay11Part1-16         	      55	  22480883 ns/op	   28337 B/op	     216 allocs/op
BenchmarkDay11Part2-16         	      44	  26099989 ns/op	   28347 B/op	     216 allocs/op
BenchmarkDay12BaselineIO-16    	   46786	     25539 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay12Part1-16         	   18145	     65829 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay12Part2-16         	   18463	     64745 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay13BaselineIO-16    	   46621	     26341 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay13Part1-16         	   41862	     29444 ns/op	    5835 B/op	      13 allocs/op
BenchmarkDay13Part2-16         	   39826	     30674 ns/op	    5843 B/op	      14 allocs/op
BenchmarkDay14BaselineIO-16    	   44380	     27030 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay14Part1-16         	    6514	    174593 ns/op	   63020 B/op	    1156 allocs/op
BenchmarkDay14Part2-16         	     142	   8285008 ns/op	 5607829 B/op	    4339 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	61.268s
```
