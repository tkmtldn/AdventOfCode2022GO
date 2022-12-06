package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func main() {
	data := GetStrings("./day04/input04.txt")
	answer1 := 0
	answer2 := 0
	for _, v := range data {
		str := strings.Replace(v, ",", " ", 1)
		str = strings.Replace(str, "-", " ", -1)
		s := strings.Split(str, " ")

		ints := []int{}
		for _, v := range s {
			i, _ := strconv.Atoi(v)
			ints = append(ints, i)
		}

		if (ints[2] <= ints[0] && ints[3] >= ints[1]) || (ints[2] >= ints[0] && ints[1] >= ints[3]) {
			answer1++
		}
		if ints[2] <= ints[1] && ints[0] <= ints[3] {
			answer2++
		}
	}
	fmt.Printf("First answer: %v \n", answer1)
	fmt.Printf("Second answer: %v \n", answer2)
}
