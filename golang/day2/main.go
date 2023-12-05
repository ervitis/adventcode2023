package main

import (
	"fmt"
	"github.com/ervitis/adventcode2023/common"
	"strconv"
	"strings"
)

const Blue = 14
const Green = 13
const Red = 12

func main() {
	fmt.Println(part2())
	fmt.Println(part1())
}

func part2() int {
	idx := 0
	sc := common.GetScanner("./day2/input")
	for sc.Scan() {
		txt := sc.Text()
		games := strings.Split(txt, ":")
		gameData := strings.TrimSpace(games[1])
		groups := strings.Split(gameData, ";")
		set := make([]map[string]int, len(groups))
		for k := range groups {
			set[k] = map[string]int{
				"blue":  0,
				"green": 0,
				"red":   0,
			}
		}

		for i, gr := range groups {
			gr = strings.TrimSpace(gr)
			balls := strings.Split(gr, ",")
			for _, ball := range balls {
				ball = strings.TrimSpace(ball)
				num, _ := strconv.Atoi(strings.Split(ball, " ")[0])
				if strings.Contains(ball, "red") {
					set[i]["red"] += num
				}
				if strings.Contains(ball, "blue") {
					set[i]["blue"] += num
				}
				if strings.Contains(ball, "green") {
					set[i]["green"] += num
				}
			}
		}
		maxB := 1
		maxG := 1
		maxR := 1
		for i := range set {
			maxB = max(maxB, set[i]["blue"])
			maxG = max(maxG, set[i]["green"])
			maxR = max(maxR, set[i]["red"])
		}
		idx += maxB * maxG * maxR
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return idx
}

func part1() int {
	idx := 0
	sc := common.GetScanner("./day2/input")
	for sc.Scan() {
		txt := sc.Text()
		games := strings.Split(txt, ":")
		gameData := strings.TrimSpace(games[1])
		groups := strings.Split(gameData, ";")
		set := make([]map[string]int, len(groups))
		for k := range groups {
			set[k] = map[string]int{
				"blue":  0,
				"green": 0,
				"red":   0,
			}
		}

		for i, gr := range groups {
			gr = strings.TrimSpace(gr)
			balls := strings.Split(gr, ",")
			for _, ball := range balls {
				ball = strings.TrimSpace(ball)
				num, _ := strconv.Atoi(strings.Split(ball, " ")[0])
				if strings.Contains(ball, "red") {
					set[i]["red"] += num
				}
				if strings.Contains(ball, "blue") {
					set[i]["blue"] += num
				}
				if strings.Contains(ball, "green") {
					set[i]["green"] += num
				}
			}
		}
		ok := true
		for i := range set {
			if set[i]["red"] > Red || set[i]["green"] > Green || set[i]["blue"] > Blue {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		game, _ := strconv.Atoi(strings.Split(games[0], " ")[1])
		idx += game
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return idx
}
