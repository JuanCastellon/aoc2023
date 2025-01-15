package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, ferr := os.Open("input.txt")
	if ferr != nil {
		panic(ferr)
	}
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile("[0-9]")
		nums := re.FindAllString(line, -1)

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("Error", err)
		} else {
			// fmt.Printf("NumZero: %v\n", numZero)
		}
		num2, err := strconv.Atoi(nums[len(nums)-1])
		if err != nil {
			fmt.Println("Error", err)
		} else {
			// fmt.Printf("NumOne: %v\n", numOne)
		}
		sum += num1*10 + num2
	}
	fmt.Printf("Value %v\n", sum)
}
