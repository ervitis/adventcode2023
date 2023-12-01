package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	f, err := os.Open("./day1/input")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()
	//scanner1 := bufio.NewScanner(f)
	scanner2 := bufio.NewScanner(f)
	//part1(scanner1)
	part2(scanner2)
}

type dataset struct {
	index        int
	numberLetter string
}
type datasets []dataset

func (d datasets) Less(i, j int) bool { return d[i].index < d[j].index }
func (d datasets) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d datasets) Len() int           { return len(d) }

func part2(scanner *bufio.Scanner) {
	set := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	findAllIndex := func(s, substr string) datasets {
		st := 0
		var d datasets
		for {
			idx := strings.Index(s[st:], substr)
			if idx == -1 {
				return d
			}
			ti := st + idx
			d = append(d, dataset{index: ti, numberLetter: set[substr]})
			st = ti + len(substr)
		}
	}

	n := 0
	var i int
	var err error

	for scanner.Scan() {
		in := make(datasets, 0)
		txt := scanner.Text()
		for k := range set {
			tin := findAllIndex(txt, k)
			if len(tin) > 0 {
				in = append(in, tin...)
			}
		}

		s := make(datasets, 0)
		for k, run := range txt {
			if unicode.IsDigit(run) {
				s = append(s, dataset{index: k, numberLetter: string(run)})
			}
		}
		data := append(in, s...)
		sort.Sort(data)

		if len(data) > 1 {
			st := fmt.Sprintf("%s%s", data[0].numberLetter, data[len(data)-1].numberLetter)
			i, err = strconv.Atoi(st)
			if err != nil {
				panic(err)
			}
		} else if len(data) == 1 {
			st := fmt.Sprintf("%s%s", data[0].numberLetter, data[0].numberLetter)
			i, err = strconv.Atoi(st)
			if err != nil {
				panic(err)
			}
		} else {
			i = 0
		}
		n += i
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func part1(scanner *bufio.Scanner) {
	regx, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}

	n := 0
	var i int
	for scanner.Scan() {
		data := strings.Join(regx.FindAllString(scanner.Text(), -1), "")
		if len(data) > 1 {
			s := fmt.Sprintf("%s%s", string(data[0]), string(data[len(data)-1]))
			i, err = strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
		} else if len(data) == 1 {
			s := fmt.Sprintf("%s%s", string(data[0]), string(data[0]))
			i, err = strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
		} else {
			i = 0
		}
		n += i
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(n)
}
