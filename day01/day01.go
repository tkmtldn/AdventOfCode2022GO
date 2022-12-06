package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
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
	data := GetStrings("./input01.txt")
	maxCalories := 0
	currentCalories := 0
	allCals := []int{}
	for _, v := range data {
		v_int, err := strconv.Atoi(v)
		if err != nil {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			allCals = append(allCals, currentCalories)
			currentCalories = 0
		} else {
			currentCalories += v_int
		}
	}

	//summing up the last elf
	allCals = append(allCals, currentCalories)
	if currentCalories > maxCalories {
		maxCalories = currentCalories
	}

	fmt.Printf("First answer: %v \n", maxCalories)

	sort.Ints(allCals)
	topThree := allCals[len(allCals)-1] + allCals[len(allCals)-2] + allCals[len(allCals)-3]
	fmt.Printf("Second answer: %v \n", topThree)
}
