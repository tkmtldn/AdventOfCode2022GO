package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetString(path string) string {
	file, err := os.Open(path)
	check(err)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scan := bufio.NewScanner(file)
	scan.Scan()
	result := scan.Text()

	check(scan.Err())

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func packetScan(s string) int {
	for i := 0; i < len(s)-5; i++ {
		line := s[i : i+4]
		if (line[0] != line[1]) && (line[0] != line[2]) && (line[0] != line[3]) && (line[1] != line[2]) && (line[1] != line[3]) && (line[2] != line[3]) {
			return i + 4
		}
	}
	return 0
}

func messageScan(s string) int {
	for i := 0; i < len(s)-15; i++ {
		line := s[i : i+14]

		total := 0
		for j := 0; j < 14; j++ {
			total += strings.Count(line, string(line[j]))
		}
		if total == 14 {
			return i + 14
		}

	}
	return 0
}

func main() {
	data := GetString("./day06/input06.txt")

	fmt.Printf("First answer: %v \n", packetScan(data))
	fmt.Printf("Second answer: %v \n", messageScan(data))
}
