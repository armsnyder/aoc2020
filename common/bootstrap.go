package common

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func Bootstrap(day int, puzzleInput interface{}) (partB bool) {
	flag.BoolVar(&partB, "b", false, "part B")
	flag.Parse()

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

	puzzlePartsRaw := bytes.Split(puzzleInputRaw, []byte{delimiter})

	switch x := puzzleInput.(type) {
	case *[]int:
		*x = make([]int, len(puzzleInputRaw))
		for i, part := range puzzlePartsRaw {
			num, err := strconv.Atoi(string(part))
			if err != nil {
				panic(err)
			}
			(*x)[i] = num
		}
	case *[]string:
		*x = make([]string, len(puzzleInputRaw))
		for i, part := range puzzlePartsRaw {
			(*x)[i] = string(part)
		}
	default:
		panic("not implemented")
	}

	return
}
