package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetData(path string) ([]string, [][]string) {
	file, err := os.Open(path)
	check(err)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scan := bufio.NewScanner(file)

	var result2 [][]string
	result3 := make([]string, 3)
	isProcedure := false

	for scan.Scan() {
		if scan.Text() == "" {
			isProcedure = true
		} else if isProcedure {
			str := strings.Replace(scan.Text(), "move", "", -1)
			str = strings.Replace(str, "from", "", -1)
			str = strings.Replace(str, "to", "", -1)
			d := strings.Split(str, " ")
			result2 = append(result2, d)
		} else {
			line := scan.Text()
			//fmt.Println(line)
			//if line[len(line)-1] == 93 {
			newLine := strings.Replace(line, "   ", "[_]", -1)
			newLine = strings.Replace(newLine, "[", "", -1)
			newLine = strings.Replace(newLine, "]", "", -1)
			newLine = strings.Replace(newLine, " ", "", -1)
			fmt.Printf(line, "-", newLine)
			for k, v := range newLine {
				if v != 95 {
					result3[k] += string(v)
				}
				//}
			}
		}

	}

	check(scan.Err())

	return result3, result2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

func beginProcs(s []string, p [][]string) string {
	for _, v := range p {
		move, _ := strconv.Atoi(v[1])
		from, _ := strconv.Atoi(v[3])
		from -= 1
		to, _ := strconv.Atoi(v[5])
		to -= 1

		cut := reverse(s[from][:move])
		uncut := s[from][move:]

		s[to] = cut + s[to]
		s[from] = uncut

	}

	ans := ""
	fmt.Println(len(s))
	for _, v := range s {
		ans += string(v[0])
	}

	return ans
}

func main() {
	stacks, procs := GetData("./day05/input_ex.txt")

	fmt.Printf("First answer: %v \n", beginProcs(stacks, procs))
	//fmt.Printf("Second answer: %v \n", data)
}
