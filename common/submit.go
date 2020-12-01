package common

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func Submit(day int, partB bool, v interface{}) {
	answer := fmt.Sprint(v)
	fmt.Println("Answer:", answer)

	f, err := os.OpenFile("solutions.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	level := "1"
	if partB {
		level = "2"
	}
	key := fmt.Sprintf("%02d.%s=", day, level)

	scanner := bufio.NewScanner(f)
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

	io.Copy(os.Stderr, resp.Body)

	if resp.StatusCode == http.StatusOK {
		if _, err := f.WriteString(key + answer + "\n"); err != nil {
			panic(err)
		}
	}
}
