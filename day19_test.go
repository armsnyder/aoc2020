package main

import (
	"reflect"
	"testing"
)

func TestDay19Part1(t *testing.T) {
	runDayTests(t, 19, []dayTest{
		{
			input: `
0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb
`,
			want: 2,
		},
	})
}

func TestDay19Part2(t *testing.T) {
	runDayTests(t, 19, []dayTest{
		{
			part2: true,
			input: `
42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba
`,
			want: 12,
		},
	})
}

func Test_day19RuleSet_match(t *testing.T) {
	type args struct {
		s        string
		rid      int
		depth    int
		maxDepth int
	}
	tests := []struct {
		name string
		r    day19RuleSet
		args args
		want []int
	}{
		{
			name: "one char ok",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{1}}},
				1: day19Rule{literal: "a"},
			},
			args: args{
				s:        "a",
				maxDepth: 10,
			},
			want: []int{1},
		},
		{
			name: "one char bad",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{1}}},
				1: day19Rule{literal: "a"},
			},
			args: args{
				s:        "b",
				maxDepth: 10,
			},
		},
		{
			name: "one other char bad",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{1}}},
				1: day19Rule{literal: "b"},
			},
			args: args{
				s:        "a",
				maxDepth: 10,
			},
		},
		{
			name: "two chars ok",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{1, 2}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
			},
			args: args{
				s:        "ab",
				maxDepth: 10,
			},
			want: []int{2},
		},
		{
			name: "two chars bad",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{1, 2}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
			},
			args: args{
				s:        "ba",
				maxDepth: 10,
			},
		},
		{
			name: "direct indirect ok",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{1, 3}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
				3: day19Rule{children: [][]int{{1, 2}}},
			},
			args: args{
				s:        "aab",
				maxDepth: 10,
			},
			want: []int{3},
		},
		{
			name: "direct indirect bad",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{1, 3}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
				3: day19Rule{children: [][]int{{1, 2}}},
			},
			args: args{
				s:        "abb",
				maxDepth: 10,
			},
		},
		{
			name: "indirect direct ok",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{3, 1}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
				3: day19Rule{children: [][]int{{1, 2}}},
			},
			args: args{
				s:        "aba",
				maxDepth: 10,
			},
			want: []int{3},
		},
		{
			name: "indirect direct bad",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{3, 1}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
				3: day19Rule{children: [][]int{{1, 2}}},
			},
			args: args{
				s:        "abb",
				maxDepth: 10,
			},
		},
		{
			name: "or directs true",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{1}, {2}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
			},
			args: args{
				s:        "a",
				maxDepth: 10,
			},
			want: []int{1},
		},
		{
			name: "or directs other true",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{1}, {2}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
			},
			args: args{
				s:        "b",
				maxDepth: 10,
			},
			want: []int{1},
		},
		{
			name: "loop ok",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{3, 2}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
				3: day19Rule{children: [][]int{{1}, {3, 1}}},
			},
			args: args{
				s:        "ab",
				maxDepth: 10,
			},
			want: []int{2},
		},
		{
			name: "loop ok with max depth",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{3, 2}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
				3: day19Rule{children: [][]int{{1}, {3, 1}}},
			},
			args: args{
				s:        "ab",
				maxDepth: 2,
			},
			want: []int{2},
		},
		{
			name: "loop ok other",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{3, 2}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
				3: day19Rule{children: [][]int{{1}, {3, 1}}},
			},
			args: args{
				s:        "aaaaab",
				maxDepth: 10,
			},
			want: []int{6},
		},
		{
			name: "loop ok other with max depth",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{3, 2}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
				3: day19Rule{children: [][]int{{1}, {3, 1}}},
			},
			args: args{
				s:        "aaaaab",
				maxDepth: 6,
			},
			want: []int{6},
		},
		{
			name: "loop ok crazy with max depth",
			r: day19RuleSet{
				0: day19Rule{children: [][]int{{3}}},
				1: day19Rule{literal: "a"},
				2: day19Rule{literal: "b"},
				3: day19Rule{children: [][]int{{1, 2}, {1, 3, 2}}},
			},
			args: args{
				s:        "aaaaaaabbbbbbb",
				maxDepth: 14,
			},
			want: []int{14},
		},
		{
			name: "loop ok crazy with max depth 2",
			r: day19RuleSet{
				0:  day19Rule{children: [][]int{{2, 8, 11, 2, 2}}},
				8:  day19Rule{children: [][]int{{42}, {42, 8}}},
				11: day19Rule{children: [][]int{{42, 31}, {42, 11, 31}}},
				42: day19Rule{children: [][]int{{1, 1, 2}}},
				31: day19Rule{children: [][]int{{2, 1, 2}}},
				1:  day19Rule{literal: "a"},
				2:  day19Rule{literal: "b"},
			},
			args: args{
				s:        "baabaabaabaabaabaabaabaabbabbabbabbb",
				maxDepth: 36,
			},
			want: []int{36},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.match(tt.args.s, tt.args.rid, tt.args.depth, tt.args.maxDepth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}
