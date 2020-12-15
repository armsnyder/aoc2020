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
BenchmarkDay01BaselineIO-16    	   48678	     25549 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay01Part1-16         	   33044	     35580 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay01Part2-16         	   31218	     38508 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay02BaselineIO-16    	   42670	     28777 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay02Part1-16         	    4366	    266745 ns/op	   95182 B/op	    2009 allocs/op
BenchmarkDay02Part2-16         	    5056	    254210 ns/op	   95181 B/op	    2009 allocs/op
BenchmarkDay03BaselineIO-16    	   45382	     26915 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay03Part1-16         	   23890	     50725 ns/op	   31186 B/op	     341 allocs/op
BenchmarkDay03Part2-16         	   20186	     58823 ns/op	   31194 B/op	     342 allocs/op
BenchmarkDay04BaselineIO-16    	   43164	     27987 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay04Part1-16         	    1927	    595645 ns/op	  271772 B/op	    5347 allocs/op
BenchmarkDay04Part2-16         	    1615	    715626 ns/op	  297500 B/op	    5879 allocs/op
BenchmarkDay05BaselineIO-16    	   44191	     26883 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay05Part1-16         	   10000	    103704 ns/op	   18626 B/op	     893 allocs/op
BenchmarkDay05Part2-16         	    6794	    184180 ns/op	   35036 B/op	     905 allocs/op
BenchmarkDay06BaselineIO-16    	   42022	     28213 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay06Part1-16         	    1279	    907314 ns/op	  330918 B/op	    4775 allocs/op
BenchmarkDay06Part2-16         	    1304	    910187 ns/op	  330875 B/op	    4774 allocs/op
BenchmarkDay07BaselineIO-16    	   36997	     31771 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay07Part1-16         	     595	   1963753 ns/op	  777855 B/op	    8538 allocs/op
BenchmarkDay07Part2-16         	     675	   1797433 ns/op	  691461 B/op	    7412 allocs/op
BenchmarkDay08BaselineIO-16    	   43204	     25766 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay08Part1-16         	   10000	    115937 ns/op	   79452 B/op	    1279 allocs/op
BenchmarkDay08Part2-16         	    5826	    214517 ns/op	  167185 B/op	    1416 allocs/op
BenchmarkDay09BaselineIO-16    	   45007	     26629 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay09Part1-16         	   10000	    106901 ns/op	   31073 B/op	    1011 allocs/op
BenchmarkDay09Part2-16         	    2958	    419719 ns/op	   31073 B/op	    1011 allocs/op
BenchmarkDay10BaselineIO-16    	   47564	     25347 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay10Part1-16         	   32599	     36999 ns/op	    6844 B/op	     119 allocs/op
BenchmarkDay10Part2-16         	   32216	     37433 ns/op	    7740 B/op	     120 allocs/op
BenchmarkDay11BaselineIO-16    	   44624	     26632 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay11Part1-16         	      52	  22078522 ns/op	   32339 B/op	     216 allocs/op
BenchmarkDay11Part2-16         	      46	  26051208 ns/op	   32350 B/op	     216 allocs/op
BenchmarkDay12BaselineIO-16    	   48334	     25259 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay12Part1-16         	   18212	     64524 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay12Part2-16         	   19168	     62050 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay13BaselineIO-16    	   48184	     25345 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay13Part1-16         	   39414	     29117 ns/op	    5835 B/op	      13 allocs/op
BenchmarkDay13Part2-16         	   40794	     29682 ns/op	    5843 B/op	      14 allocs/op
BenchmarkDay14BaselineIO-16    	   44528	     27191 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay14Part1-16         	    7152	    173005 ns/op	   63014 B/op	    1155 allocs/op
BenchmarkDay14Part2-16         	     144	   7962902 ns/op	 5606082 B/op	    4330 allocs/op
BenchmarkDay15BaselineIO-16    	   48751	     25069 ns/op	     370 B/op	       7 allocs/op
BenchmarkDay15Part1-16         	   12550	     95303 ns/op	   26648 B/op	      46 allocs/op
BenchmarkDay15Part2-16         	       1	2120114370 ns/op	352250008 B/op	  153002 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	65.572s
```
