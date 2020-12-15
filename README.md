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
BenchmarkDay01Part1-16    	   32583	     36044 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay01Part2-16    	   30999	     40010 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay02Part1-16    	    4491	    261355 ns/op	   95181 B/op	    2009 allocs/op
BenchmarkDay02Part2-16    	    5140	    254423 ns/op	   95181 B/op	    2009 allocs/op
BenchmarkDay03Part1-16    	   23983	     49793 ns/op	   31186 B/op	     341 allocs/op
BenchmarkDay03Part2-16    	   19681	     58537 ns/op	   31194 B/op	     342 allocs/op
BenchmarkDay04Part1-16    	    1972	    601475 ns/op	  271772 B/op	    5347 allocs/op
BenchmarkDay04Part2-16    	    1695	    715204 ns/op	  297488 B/op	    5879 allocs/op
BenchmarkDay05Part1-16    	   10000	    103739 ns/op	   18626 B/op	     893 allocs/op
BenchmarkDay05Part2-16    	    6868	    185878 ns/op	   35036 B/op	     905 allocs/op
BenchmarkDay06Part1-16    	    1309	    909137 ns/op	  330949 B/op	    4776 allocs/op
BenchmarkDay06Part2-16    	    1267	    946114 ns/op	  330873 B/op	    4774 allocs/op
BenchmarkDay07Part1-16    	     585	   1949697 ns/op	  777201 B/op	    8538 allocs/op
BenchmarkDay07Part2-16    	     618	   1806132 ns/op	  691021 B/op	    7412 allocs/op
BenchmarkDay08Part1-16    	   10000	    122598 ns/op	   79453 B/op	    1279 allocs/op
BenchmarkDay08Part2-16    	    5670	    219290 ns/op	  167186 B/op	    1416 allocs/op
BenchmarkDay09Part1-16    	   10000	    105901 ns/op	   31073 B/op	    1011 allocs/op
BenchmarkDay09Part2-16    	    2734	    420242 ns/op	   31074 B/op	    1011 allocs/op
BenchmarkDay10Part1-16    	   31514	     36921 ns/op	    6844 B/op	     119 allocs/op
BenchmarkDay10Part2-16    	   32199	     37822 ns/op	    7740 B/op	     120 allocs/op
BenchmarkDay11Part1-16    	      51	  23592200 ns/op	   32340 B/op	     216 allocs/op
BenchmarkDay11Part2-16    	      45	  26101968 ns/op	   32346 B/op	     216 allocs/op
BenchmarkDay12Part1-16    	   17924	     65021 ns/op	    6811 B/op	     799 allocs/op
BenchmarkDay12Part2-16    	   18420	     62887 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay13Part1-16    	   40755	     28892 ns/op	    5835 B/op	      13 allocs/op
BenchmarkDay13Part2-16    	   38588	     30400 ns/op	    5843 B/op	      14 allocs/op
BenchmarkDay14Part1-16    	    7243	    172680 ns/op	   63017 B/op	    1155 allocs/op
BenchmarkDay14Part2-16    	     145	   8220629 ns/op	 5608410 B/op	    4339 allocs/op
BenchmarkDay15Part1-16    	   36793	     32832 ns/op	   12847 B/op	      14 allocs/op
BenchmarkDay15Part2-16    	       2	 516798013 ns/op	120011536 B/op	      18 allocs/op
BenchmarkBaselineIO-16    	   44518	     25628 ns/op	     370 B/op	       7 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	44.243s
```
