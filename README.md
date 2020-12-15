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
The `BenchmarkBaselineIO` benchmark reads and discards the day 1 input file without running the
solution code, to serve as a baseline for the solution benchmarks.

```
$ go test -run=XXX -bench=.
goos: darwin
goarch: amd64
pkg: github.com/armsnyder/aoc2020
BenchmarkDay01Part1-16    	   31020	     36692 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay01Part2-16    	   29131	     40127 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay02Part1-16    	    3925	    267512 ns/op	   95180 B/op	    2009 allocs/op
BenchmarkDay02Part2-16    	    4684	    252246 ns/op	   95181 B/op	    2009 allocs/op
BenchmarkDay03Part1-16    	   23142	     51751 ns/op	   31186 B/op	     341 allocs/op
BenchmarkDay03Part2-16    	   19556	     62405 ns/op	   31194 B/op	     342 allocs/op
BenchmarkDay04Part1-16    	    1981	    618916 ns/op	  271771 B/op	    5347 allocs/op
BenchmarkDay04Part2-16    	    1578	    711282 ns/op	  297507 B/op	    5879 allocs/op
BenchmarkDay05Part1-16    	   10000	    105318 ns/op	   18626 B/op	     893 allocs/op
BenchmarkDay05Part2-16    	    6783	    187736 ns/op	   35036 B/op	     905 allocs/op
BenchmarkDay06Part1-16    	    1293	    921388 ns/op	  330884 B/op	    4775 allocs/op
BenchmarkDay06Part2-16    	    1214	    942798 ns/op	  331000 B/op	    4777 allocs/op
BenchmarkDay07Part1-16    	     600	   2051809 ns/op	  777836 B/op	    8538 allocs/op
BenchmarkDay07Part2-16    	     639	   1853083 ns/op	  691761 B/op	    7412 allocs/op
BenchmarkDay08Part1-16    	   10000	    117075 ns/op	   79453 B/op	    1279 allocs/op
BenchmarkDay08Part2-16    	    5569	    218673 ns/op	  167186 B/op	    1416 allocs/op
BenchmarkDay09Part1-16    	   10000	    112299 ns/op	   31073 B/op	    1011 allocs/op
BenchmarkDay09Part2-16    	    2881	    427526 ns/op	   31074 B/op	    1011 allocs/op
BenchmarkDay10Part1-16    	   32011	     37485 ns/op	    6844 B/op	     119 allocs/op
BenchmarkDay10Part2-16    	   31807	     37891 ns/op	    7740 B/op	     120 allocs/op
BenchmarkDay11Part1-16    	      51	  22589550 ns/op	   32340 B/op	     216 allocs/op
BenchmarkDay11Part2-16    	      46	  27156180 ns/op	   32345 B/op	     216 allocs/op
BenchmarkDay12Part1-16    	   17610	     64810 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay12Part2-16    	   18345	     64872 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay13Part1-16    	   39361	     29788 ns/op	    5835 B/op	      13 allocs/op
BenchmarkDay13Part2-16    	   37503	     30035 ns/op	    5843 B/op	      14 allocs/op
BenchmarkDay14Part1-16    	    6602	    178433 ns/op	   63032 B/op	    1156 allocs/op
BenchmarkDay14Part2-16    	     141	   8589890 ns/op	 5607276 B/op	    4335 allocs/op
BenchmarkDay15Part1-16    	   35192	     34794 ns/op	   21044 B/op	      14 allocs/op
BenchmarkDay15Part2-16    	       2	 671180872 ns/op	240008000 B/op	      18 allocs/op
BenchmarkBaselineIO-16    	   43400	     25269 ns/op	     369 B/op	       7 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	44.716s
```
