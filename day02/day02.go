package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetStrings(path string) []string {
	file, err := os.Open(path)
	check(err)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scan := bufio.NewScanner(file)

	var result []string

	for scan.Scan() {
		result = append(result, scan.Text())
	}

	check(scan.Err())

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func game1(opp int, you int) int {
	you -= 87
	opp -= 64
	if you == opp {
		return 3 + you
	} else if (you-opp == 1) || (you-opp == -2) {
		return 6 + you
	} else {
		return 0 + you
	}
}

func game2(opp int, you int) int {
	you -= 87
	opp -= 64
	if you == 2 {
		return 3 + opp
	} else if you == 1 {
		res := opp - 1
		if res == 0 {
			return 0 + 3
		}
		return 0 + opp - 1
	} else {
		res := opp + 1
		if res == 4 {
			return 6 + 1
		}
		return 6 + (opp + 1)
	}
}

func main() {
	data := GetStrings("./day02/input02.txt")

	totalScore := 0
	totalScore2 := 0
	for _, v := range data {
		totalScore += game1(int(v[0]), int(v[2]))
		totalScore2 += game2(int(v[0]), int(v[2]))
	}
	fmt.Printf("First answer: %v \n", totalScore)
	fmt.Printf("Second answer: %v \n", totalScore2)
}
