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
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/armsnyder/aoc2020
BenchmarkBaselineIO-16    	   44950	     25206 ns/op	     369 B/op	       7 allocs/op
BenchmarkDay/01/Part1-16  	   32905	     36276 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay/01/Part2-16  	   28926	     39114 ns/op	    9373 B/op	     218 allocs/op
BenchmarkDay/02/Part1-16  	    4392	    251253 ns/op	   95181 B/op	    2009 allocs/op
BenchmarkDay/02/Part2-16  	    5166	    245086 ns/op	   95181 B/op	    2009 allocs/op
BenchmarkDay/03/Part1-16  	   23556	     49908 ns/op	   31186 B/op	     341 allocs/op
BenchmarkDay/03/Part2-16  	   20019	     59153 ns/op	   31194 B/op	     342 allocs/op
BenchmarkDay/04/Part1-16  	    2030	    593450 ns/op	  271770 B/op	    5347 allocs/op
BenchmarkDay/04/Part2-16  	    1641	    709610 ns/op	  297471 B/op	    5878 allocs/op
BenchmarkDay/05/Part1-16  	   10000	    103173 ns/op	   18626 B/op	     893 allocs/op
BenchmarkDay/05/Part2-16  	    6789	    182711 ns/op	   35035 B/op	     905 allocs/op
BenchmarkDay/06/Part1-16  	    1254	    883501 ns/op	  330940 B/op	    4776 allocs/op
BenchmarkDay/06/Part2-16  	    1306	    905319 ns/op	  330905 B/op	    4775 allocs/op
BenchmarkDay/07/Part1-16  	     613	   1978018 ns/op	  776738 B/op	    8538 allocs/op
BenchmarkDay/07/Part2-16  	     651	   1799740 ns/op	  690940 B/op	    7412 allocs/op
BenchmarkDay/08/Part1-16  	   10000	    119701 ns/op	   79453 B/op	    1279 allocs/op
BenchmarkDay/08/Part2-16  	    4071	    252675 ns/op	  167185 B/op	    1416 allocs/op
BenchmarkDay/09/Part1-16  	   10000	    107844 ns/op	   31073 B/op	    1011 allocs/op
BenchmarkDay/09/Part2-16  	    2923	    416603 ns/op	   31073 B/op	    1011 allocs/op
BenchmarkDay/10/Part1-16  	   33099	     36761 ns/op	    6844 B/op	     119 allocs/op
BenchmarkDay/10/Part2-16  	   28045	     37273 ns/op	    7740 B/op	     120 allocs/op
BenchmarkDay/11/Part1-16  	      55	  21846868 ns/op	   32337 B/op	     216 allocs/op
BenchmarkDay/11/Part2-16  	      45	  25763523 ns/op	   32351 B/op	     216 allocs/op
BenchmarkDay/12/Part1-16  	   18768	     65053 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay/12/Part2-16  	   19299	     61873 ns/op	    6812 B/op	     799 allocs/op
BenchmarkDay/13/Part1-16  	   41491	     29056 ns/op	    5835 B/op	      13 allocs/op
BenchmarkDay/13/Part2-16  	   40671	     29303 ns/op	    5843 B/op	      14 allocs/op
BenchmarkDay/14/Part1-16  	    6981	    170990 ns/op	   63015 B/op	    1155 allocs/op
BenchmarkDay/14/Part2-16  	     148	   8191668 ns/op	 5610704 B/op	    4347 allocs/op
BenchmarkDay/15/Part1-16  	   37526	     32416 ns/op	   12847 B/op	      14 allocs/op
BenchmarkDay/15/Part2-16  	       3	 426698431 ns/op	120011461 B/op	      17 allocs/op
PASS
ok  	github.com/armsnyder/aoc2020	45.496s
```
