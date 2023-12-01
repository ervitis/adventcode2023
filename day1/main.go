package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./day1/input")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = f.Close()
	}()
	scanner := bufio.NewScanner(f)
	part1(scanner)
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
