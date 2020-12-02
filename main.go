package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type dayFn func(part2 bool, inputUntyped interface{}) interface{}

var days = make(map[int]dayFn)

func declareDay(day int, dayFn dayFn) interface{} {
	days[day] = dayFn
	return nil
}

func main() {
	easternTime, err := time.LoadLocation("EST")
	if err != nil {
		panic(err)
	}

	var (
		day   int
		part2 bool
	)
	flag.IntVar(&day, "d", time.Now().In(easternTime).Day(), "Day of the month")
	flag.BoolVar(&part2, "2", false, "Run part 2")
	flag.Parse()

	dayFn, ok := days[day]
	if !ok {
		fmt.Printf("Generating a stub for day %d\n", day)
		generateStub(day)
		return
	}

	input := getInput(day)

	start := time.Now()
	output := dayFn(part2, input)
	runTime := time.Since(start)
	fmt.Printf("Finished in %s\n", runTime)

	submit(day, part2, output)
}

func getInput(day int) interface{} {
	if err := os.MkdirAll("inputs", 0755); err != nil {
		panic(err)
	}

	puzzleInputFilename := filepath.Join("inputs", fmt.Sprintf("%02d.txt", day))

	if _, err := os.Lstat(puzzleInputFilename); errors.Is(err, os.ErrNotExist) {
		session, err := ioutil.ReadFile("session.txt")
		if err != nil {
			panic(err)
		}
		session = bytes.TrimSpace(session)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2020/day/%d/input", day), nil)
		if err != nil {
			panic(err)
		}

		req.AddCookie(&http.Cookie{
			Name:  "session",
			Value: string(session),
		})

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			panic(resp.Status)
		}

		out, err := os.Create(puzzleInputFilename)
		if err != nil {
			panic(err)
		}
		defer out.Close()

		if _, err := io.Copy(out, resp.Body); err != nil {
			panic(err)
		}
	}

	puzzleInputRaw, err := ioutil.ReadFile(puzzleInputFilename)
	if err != nil {
		panic(err)
	}

	puzzleInputRaw = bytes.TrimSpace(puzzleInputRaw)

	var delimiter byte
	switch {
	case bytes.ContainsRune(puzzleInputRaw, '\n'):
		delimiter = '\n'
	case bytes.ContainsRune(puzzleInputRaw, ','):
		delimiter = ','
	default:
		delimiter = '\n'
	}

	puzzleSplit := bytes.Split(puzzleInputRaw, []byte{delimiter})

	if _, err := strconv.Atoi(string(puzzleSplit[0])); err == nil {
		input := make([]int, len(puzzleSplit))
		for i, item := range puzzleSplit {
			part, err := strconv.Atoi(string(item))
			if err != nil {
				panic(err)
			}
			input[i] = part
		}
		return input
	}

	input := make([]string, len(puzzleSplit))
	for i, item := range puzzleSplit {
		input[i] = string(item)
	}
	return input
}

func submit(day int, part2 bool, v interface{}) {
	answer := fmt.Sprint(v)
	fmt.Println("Answer:", answer)

	solutions, err := os.OpenFile("solutions.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer solutions.Close()

	level := "1"
	if part2 {
		level = "2"
	}
	key := fmt.Sprintf("%02d.%s=", day, level)

	scanner := bufio.NewScanner(solutions)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), key) {
			correctAnswer := strings.TrimPrefix(scanner.Text(), key)
			if answer == correctAnswer {
				fmt.Print("Correct")
			} else {
				fmt.Print("Incorrect")
			}
			fmt.Println(" (Already submitted)")
			return
		}
	}

	fmt.Println("Submitting...")

	session, err := ioutil.ReadFile("session.txt")
	if err != nil {
		panic(err)
	}
	session = bytes.TrimSpace(session)

	values := make(url.Values)
	values.Set("answer", answer)
	values.Set("level", level)

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("https://adventofcode.com/2020/day/%d/answer", day),
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		panic(err)
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: string(session),
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if _, err := io.Copy(os.Stderr, resp.Body); err != nil {
			panic(err)
		}
		panic(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	message := string(regexp.MustCompile("<p[^>]*>(.*)</p>").FindSubmatch(body)[1])
	fmt.Println(message)

	if strings.Contains(message, "That's the right answer!") {
		if _, err := solutions.WriteString(key + answer + "\n"); err != nil {
			panic(err)
		}
	}
}

func generateStub(day int) {
	code, err := os.Create(fmt.Sprintf("day%02d.go", day))
	if err != nil {
		panic(err)
	}
	defer code.Close()

	test, err := os.Create(fmt.Sprintf("day%02d_test.go", day))
	if err != nil {
		panic(err)
	}
	defer code.Close()

	if _, err := fmt.Fprintf(code, `package main

var _ = declareDay(%[1]d, day%02[1]d)

func day%02[1]d(part2 bool, inputUntyped interface{}) interface{} {
	panic("no solution")
}
`, day); err != nil {
		panic(err)
	}

	if _, err := fmt.Fprintf(test, `package main

import (
	"reflect"
	"testing"
)

func Test_day%02[1]d(t *testing.T) {
	type args struct {
		part2        bool
		inputUntyped interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day%02[1]d(tt.args.part2, tt.args.inputUntyped); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("day%02[1]d(%[2]s, ...) = %[2]s, want %[2]s", tt.args.part2, got, tt.want)
			}
		})
	}
}
`, day, "%v"); err != nil {
		panic(err)
	}
}