package main

import(
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

func convert(x string) int {
	num, err := strconv.Atoi(x)
	
	if err != nil {
		fmt.Println("Error converting string:", err)		
	}
	return num
}

func main() {
	inputStr := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	splitStr := strings.Split(inputStr,": ")
	
	//Regex for getting numbers from a string
	reN := regexp.MustCompile("[0-9]+")
	
	//Getting the card number
	cardNumStr := reN.FindAllString(splitStr[0],-1)
	cardNum := convert(cardNumStr[0])
	
	//Card strings
	cardStr := strings.Split(splitStr[1],"|")
	myNumStr := cardStr[1]
	winNumStr := cardStr[0]
	
	//Getting the arrays of numbers
	myNums := reN.FindAllString(myNumStr,-1)
	winNums := reN.FindAllString(winNumStr,-1)
	
	//First flag for determining if we multiply the score
	sum := 0
	first := true
	
	//iterate through every one of our numbers and comparing to the winning number
	for i := 0; i < len(myNums); i++{
		for j := 0; j < len(winNums); j++{
			if myNums[i] == winNums[j] && first{
				sum++
				first = false
			} else if myNums[i] == winNums[j] && !first{
				sum = sum*2
			}
		}
		first = false
	}
	fmt.Print("Score for Card #",cardNum)
	fmt.Println(": ",sum)
}
