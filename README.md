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
BenchmarkDay01BaselineIO-16        41938             27435 ns/op             594 B/op         10 allocs/op
BenchmarkDay01Part1-16             29422             39934 ns/op            9597 B/op        221 allocs/op
BenchmarkDay01Part2-16             28118             43403 ns/op            9597 B/op        221 allocs/op
BenchmarkDay02BaselineIO-16        38353             31125 ns/op             594 B/op         10 allocs/op
BenchmarkDay02Part1-16              4461            258624 ns/op           95405 B/op       2012 allocs/op
BenchmarkDay02Part2-16              4851            243559 ns/op           95405 B/op       2012 allocs/op
BenchmarkDay03BaselineIO-16        40402             29151 ns/op             594 B/op         10 allocs/op
BenchmarkDay03Part1-16             21900             54198 ns/op           31417 B/op        345 allocs/op
BenchmarkDay03Part2-16             18831             62936 ns/op           31418 B/op        345 allocs/op
BenchmarkDay04BaselineIO-16        38607             30632 ns/op             595 B/op         10 allocs/op
BenchmarkDay04Part1-16              1785            606372 ns/op          272013 B/op       5351 allocs/op
BenchmarkDay04Part2-16              1500            717397 ns/op          297716 B/op       5883 allocs/op
BenchmarkDay05BaselineIO-16        40477             29491 ns/op             594 B/op         10 allocs/op
BenchmarkDay05Part1-16             10000            112255 ns/op           18850 B/op        896 allocs/op
BenchmarkDay05Part2-16              6514            191777 ns/op           35259 B/op        908 allocs/op
BenchmarkDay06BaselineIO-16        38810             30787 ns/op             594 B/op         10 allocs/op
BenchmarkDay06Part1-16              1182            934322 ns/op          331181 B/op       4779 allocs/op
BenchmarkDay06Part2-16              1190            966327 ns/op          331178 B/op       4779 allocs/op
BenchmarkDay07BaselineIO-16        32940             35183 ns/op             595 B/op         10 allocs/op
BenchmarkDay07Part1-16               570           2024764 ns/op          778326 B/op       8542 allocs/op
BenchmarkDay07Part2-16               642           1843328 ns/op          690880 B/op       7414 allocs/op
BenchmarkDay08BaselineIO-16        42567             28103 ns/op             594 B/op         10 allocs/op
BenchmarkDay08Part1-16              9686            122132 ns/op           79677 B/op       1282 allocs/op
BenchmarkDay08Part2-16              4902            232660 ns/op          167409 B/op       1419 allocs/op
BenchmarkDay09BaselineIO-16        40386             30621 ns/op             595 B/op         10 allocs/op
BenchmarkDay09Part1-16              8821            138360 ns/op           31297 B/op       1014 allocs/op
BenchmarkDay09Part2-16              2523            441316 ns/op           31297 B/op       1014 allocs/op
BenchmarkDay10BaselineIO-16        41451             28041 ns/op             594 B/op         10 allocs/op
BenchmarkDay10Part1-16             28278             40586 ns/op            7068 B/op        122 allocs/op
BenchmarkDay10Part2-16             29238             39848 ns/op            7964 B/op        123 allocs/op
BenchmarkDay11BaselineIO-16        40863             30248 ns/op             595 B/op         10 allocs/op
BenchmarkDay11Part1-16                33          35356742 ns/op         5442572 B/op      58538 allocs/op
BenchmarkDay11Part2-16                25          46377738 ns/op        10083269 B/op      81993 allocs/op
BenchmarkDay12BaselineIO-16        41556             27731 ns/op             595 B/op         10 allocs/op
BenchmarkDay12Part1-16             16970             70356 ns/op            7035 B/op        802 allocs/op
BenchmarkDay12Part2-16             17464             67859 ns/op            7036 B/op        802 allocs/op
BenchmarkDay13BaselineIO-16        42915             27642 ns/op             594 B/op         10 allocs/op
BenchmarkDay13Part1-16             37333             31562 ns/op            6067 B/op         17 allocs/op
BenchmarkDay13Part2-16             35082             32941 ns/op            6067 B/op         17 allocs/op
BenchmarkDay14BaselineIO-16        38347             29509 ns/op             595 B/op         10 allocs/op
BenchmarkDay14Part1-16              5944            182072 ns/op           63246 B/op       1159 allocs/op
BenchmarkDay14Part2-16               135           8283190 ns/op         5609086 B/op       4341 allocs/op
PASS
ok      github.com/armsnyder/aoc2020    62.236s
```
