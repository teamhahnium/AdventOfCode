package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sumTarget := 2020
	pt1, err := FindMultiplyPartOne(sumTarget)

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	fmt.Println("Answer Part 1 : " + strconv.Itoa(pt1))
}

// FindMultiplyPartOne of two numbers that equals sumTarget when added.
func FindMultiplyPartOne(sumTarget int) (answer int, err error) {
	s := NewSet()
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		if target := sumTarget - i; s.Contains(target) {
			return i * target, nil
		}

		s.Add(i)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	return 0, fmt.Errorf("No answer found")
}
