package main

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
)

type stack []string

func (s stack) IsEmpty() bool {
	if len(s) == 0{
		return true	
	} else {
		return false	
	}
}

func (s stack) Push(v string) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, string) {
    stackLen := len(s)
    return  s[:stackLen-1], s[stackLen-1]
}

func convert(x string) int {
	num, err := strconv.Atoi(x)
	
	if err != nil {
		fmt.Println("Error converting string:", err)		
	}
	return num
}

func main() {
	s1 := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	s2 := "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19"
	s3 := "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1"
	s4 := "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"
	s5 := "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36"
	s6 := "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
	
	//Building our original array of scratchcards
	var orig [6]string = [6]string{s1,s2,s3,s4,s5,s6}
	
	//Regex for extracting the scratchcard number
	reN := regexp.MustCompile("[0-9]+")
	
	//Creating the stack of dupes that score and produce more dupes
	s := make(stack,0)
	
	//Counter for all the scratchcards we go through
	cardCount := 0
	
	//Processing our original set of scratchcards and making dupe array
	for i:=0;i<len(orig);i++{
		splitStr := strings.Split(orig[i],": ")
		
		//Split our new string to get winning/our numbers
		cardStr := strings.Split(splitStr[1],"|")
		myNumStr := cardStr[1]
		winNumStr := cardStr[0]
		
		//Get workable arrays of winning/our numbers
		myNums := reN.FindAllString(myNumStr,-1)
		winNums := reN.FindAllString(winNumStr,-1)
		
		counter := 0
		//Checking our number against winning numbers
		for j:=0;j<len(myNums);j++{ 
			for k:=0;k<len(winNums);k++{
				if myNums[j] == winNums[k]{
					counter++
				}
			}
		}
		//Checking to see if we had any matches
		//If statement makes sure we're in the bounds of the array
		if counter > 0 && i+counter <= 6 {
			for index:=i+1;index<=i+counter;index++{
				s = s.Push(orig[index])
			}
		//If statement triggers in the event the indexes exceed bounds
		} else if counter > 0 && i+counter > 6{
			for index:=i+1;index<=len(orig);index++{
				s = s.Push(orig[index])
			}
		//No matches found in card, nothing happens.
		}
		counter = 0
		cardCount++
	}
	
	//Pop scratchcard off the stack
	//Process card's numbers, push new cards to stack
	//Repeat until the stack is empty
	for !s.IsEmpty(){
		tempStack, cardStr := s.Pop()
		s = tempStack
		
		//Processing the contents of the cardStr
		splitStr := strings.Split(cardStr,": ")
		numberStr := reN.FindAllString(splitStr[0],-1)
		cardNumber := convert(numberStr[0])
		cardIndex := cardNumber - 1
		
		numStr := strings.Split(splitStr[1],"|")
		
		myNums := reN.FindAllString(numStr[0],-1)
		winNums := reN.FindAllString(numStr[1],-1)
		
		counter := 0
		//Checking our number against winning numbers
		for i:=0;i<len(myNums);i++{ 
			for j:=0;j<len(winNums);j++{
				if myNums[i] == winNums[j]{
					counter++
				}
			}
		}
		//Checking to see if we had any matches
		//If statement makes sure we're in the bounds of the array
		if counter > 0 && cardIndex+counter <= 6 {
			for i:=cardIndex;i<cardIndex+counter;i++{
				s = s.Push(orig[i])
			}
		//If statement triggers in the event the indexes exceed bounds
		} else if counter > 0 && cardIndex+counter > 6 {
			for i:=cardIndex;i<len(orig);i++{
				s = s.Push(orig[i])
			}
		}
		cardCount++
	}
	fmt.Println("Total Number of Scratchers: ",cardCount)
}
