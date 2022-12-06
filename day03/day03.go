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

func game(f string, s string) int {
	res := []int{}
	for _, v1 := range f {
		for _, v2 := range s {
			if v1 == v2 {
				res = append(res, int(v2))
			}
		}
	}
	ans := int(res[0]) - 38
	if ans < 53 {
		return ans
	}
	return ans - 58
}

func game2(f string, s string, t string) int {
	res := []int{}
	for _, v1 := range f {
		for _, v2 := range s {
			for _, v3 := range t {
				if (v1 == v2) && (v2 == v3) {
					res = append(res, int(v3))
				}
			}
		}
	}

	ans := int(res[0]) - 38
	if ans < 53 {
		return ans
	}
	return ans - 58
}

func main() {
	data := GetStrings("./day03/input03.txt")

	answer1 := 0

	count := 0
	secondPack := []string{}
	answer2 := 0

	for _, v := range data {
		len := len(v)
		first := v[:len/2]
		second := v[len/2:]
		answer1 += game(first, second)

		count++
		secondPack = append(secondPack, v)
		if count%3 == 0 {
			//fmt.Println(secondPack[0], secondPack[1], secondPack[2])
			answer2 += game2(secondPack[0], secondPack[1], secondPack[2])
			secondPack = []string{}
		}
	}

	fmt.Printf("First answer: %v \n", answer1)
	fmt.Printf("Second answer: %v \n", answer2)
}
