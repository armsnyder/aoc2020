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

func Test_day19RuleSet_matchRuleZero(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		ruleSet day19RuleSet
		message string
		want    bool
	}{
		{
			name: "one literal ok",
			ruleSet: day19RuleSet{
				0: day19Rule{children: [][]int{{1}}},
				1: day19Rule{literal: 'a'},
			},
			message: "a",
			want:    true,
		},
		{
			name: "one literal bad",
			ruleSet: day19RuleSet{
				0: day19Rule{children: [][]int{{1}}},
				1: day19Rule{literal: 'a'},
			},
			message: "b",
			want:    false,
		},
		{
			name: "two literals ok",
			ruleSet: day19RuleSet{
				0: day19Rule{children: [][]int{{1, 2}}},
				1: day19Rule{literal: 'a'},
				2: day19Rule{literal: 'b'},
			},
			message: "ab",
			want:    true,
		},
		{
			name: "two literals bad",
			ruleSet: day19RuleSet{
				0: day19Rule{children: [][]int{{1, 2}}},
				1: day19Rule{literal: 'a'},
				2: day19Rule{literal: 'b'},
			},
			message: "ba",
			want:    false,
		},
		{
			name: "literal then complex ok",
			ruleSet: day19RuleSet{
				0: day19Rule{children: [][]int{{1, 3}}},
				1: day19Rule{literal: 'a'},
				2: day19Rule{literal: 'b'},
				3: day19Rule{children: [][]int{{1, 2}}},
			},
			message: "aab",
			want:    true,
		},
		{
			name: "literal then complex bad",
			ruleSet: day19RuleSet{
				0: day19Rule{children: [][]int{{1, 3}}},
				1: day19Rule{literal: 'a'},
				2: day19Rule{literal: 'b'},
				3: day19Rule{children: [][]int{{1, 2}}},
			},
			message: "abb",
			want:    false,
		},
		{
			name: "complex then literal ok",
			ruleSet: day19RuleSet{
				0: day19Rule{children: [][]int{{3, 1}}},
				1: day19Rule{literal: 'a'},
				2: day19Rule{literal: 'b'},
				3: day19Rule{children: [][]int{{1, 2}}},
			},
			message: "aba",
			want:    true,
		},
		{
			name: "complex then literal bad",
			ruleSet: day19RuleSet{
				0: day19Rule{children: [][]int{{3, 1}}},
				1: day19Rule{literal: 'a'},
				2: day19Rule{literal: 'b'},
				3: day19Rule{children: [][]int{{1, 2}}},
			},
			message: "abb",
			want:    false,
		},
		{
			name: "or literals true a",
			ruleSet: day19RuleSet{
				0: day19Rule{children: [][]int{{1}, {2}}},
				1: day19Rule{literal: 'a'},
				2: day19Rule{literal: 'b'},
			},
			message: "a",
			want:    true,
		},
		{
			name: "or literals true b",
			ruleSet: day19RuleSet{
				0: day19Rule{children: [][]int{{1}, {2}}},
				1: day19Rule{literal: 'a'},
				2: day19Rule{literal: 'b'},
			},
			message: "b",
			want:    true,
		},
		{
			name: "loop ok",
			ruleSet: day19RuleSet{
				0:  day19Rule{children: [][]int{{2, 8, 11, 2, 2}}},
				8:  day19Rule{children: [][]int{{42}, {42, 8}}},
				11: day19Rule{children: [][]int{{42, 31}, {42, 11, 31}}},
				42: day19Rule{children: [][]int{{1, 1, 2}}},
				31: day19Rule{children: [][]int{{2, 1, 2}}},
				1:  day19Rule{literal: 'a'},
				2:  day19Rule{literal: 'b'},
			},
			message: "baabaabaabaabaabaabaabaabbabbabbabbb",
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ruleSet.matchRuleZero(tt.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matchRuleZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
