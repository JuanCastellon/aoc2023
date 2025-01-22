package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convert(x string) int {
	num, err := strconv.Atoi(x)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	return num
}

func main() {
	file, ferr := os.Open("input.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)
	answer := 0
	flag := -1

	for scanner.Scan() {
		//Getting line of input from file
		line := scanner.Text()

		//Splitting the game number and subgames strings
		split := strings.Split(line, ": ")
		subGameAry := strings.Split(split[1], "; ") //All the sub games we will test

		//Creating the regex necessary for getting numbers array from input
		reNA := regexp.MustCompile("[0-9]+ r|[0-9]+ g|[0-9]+ b")

		//Regex for getting a number from the number array
		reN := regexp.MustCompile("[0-9]+")

		//Getting the game number
		gameStr := reN.FindAllString(split[0], -1)
		game := convert(gameStr[0])

		//Looping through sub games, reds > 12, greens > 13, blues > 14
		for i := 0; i < len(subGameAry); i++ {
			numAry := reNA.FindAllString(subGameAry[i], -1)
			for j := 0; j < len(numAry); j++ {
				if strings.ContainsRune(numAry[j], 'r') {
					temp := reN.FindAllString(numAry[j], -1)
					number := convert(temp[0])
					if number > 12 {
						flag = 1
					}
				}
				if strings.ContainsRune(numAry[j], 'g') {
					temp := reN.FindAllString(numAry[j], -1)
					number := convert(temp[0])
					if number > 13 {
						flag = 1
					}
				}
				if strings.ContainsRune(numAry[j], 'b') {
					temp := reN.FindAllString(numAry[j], -1)
					number := convert(temp[0])
					if number > 14 {
						flag = 1
					}
				}
			}
		}

		if flag == 1 {
			fmt.Print("Game ", game)
			fmt.Println(" is not possible.")
		} else {
			fmt.Print("Game ", game)
			fmt.Println(" is possible.")
			answer += game
		}
		flag = -1
	}
	fmt.Println("Answer: ", answer)

}
