package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func convert(x string) int {
	num, err := strconv.Atoi(x)

	if err != nil {
		fmt.Println("Error: ", err)
	}
	return num
}
func main() {
	inputStr := "nineightwone"
	re := regexp.MustCompile("[0-9]|twone|oneight|nineight|one|two|three|four|five|six|seven|eight|nine")
	tokenAry := re.FindAllString(inputStr, -1)
	fmt.Println(tokenAry)
	//Getting first and last element of the token array
	fStr := tokenAry[0]
	lStr := tokenAry[len(tokenAry)-1]
	
	fNum := 0
	lNum := 0
	
	//Making sure that the first element isn't an edge case we can 			trivialize
	if fStr == "oneight"{
		fNum = 1
	} else if fStr == "twone" {
		fNum = 2
	} else if fStr == "nineight" {
		fNum = 9
	} else {
		//if fStr passes these conditionals, then it must be a digit in 		number or alphabetical form. either way still string type
		switch fStr {
			case "one":
				fNum = 1
			case "two":
				fNum = 2
			case "three":
				fNum = 3
			case "four":
				fNum = 4
			case "five":
				fNum = 5
			case "six":
				fNum = 6
			case "seven":
				fNum = 7
			case "eight":
				fNum = 8
			case "nine":
				fNum = 9
			default:
				fNum = convert(fStr)
		}
	}
	
	//Making sure the last element isn't an edge case we can trivialize
	if lStr == "oneight"{
		lNum = 8
	} else if lStr == "twone" {
		lNum = 1
	} else if lStr == "nineight" {
		lNum = 8
	} else { 
		//if lStr passes these conditionals, then it must be a digit in 		number or alphabetical form. either way still string type
		switch lStr {
			case "one":
				lNum = 1
			case "two":
				lNum = 2
			case "three":
				lNum = 3
			case "four":
				lNum = 4
			case "five":
				lNum = 5
			case "six":
				lNum = 6
			case "seven":
				lNum = 7
			case "eight":
				lNum = 8
			case "nine":
				lNum = 9
			default:
				lNum = convert(lStr)
		}	
	}
	
	caliVal := fNum*10 + lNum
	fmt.Println("Calibration Value: ", caliVal)
}
