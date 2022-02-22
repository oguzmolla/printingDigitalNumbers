package main

import "fmt"

func main() {
	var x int32 = 9
	printNumber(x)
}

func fullLine(width int32) {
	var i int32

	for i = 0; i < width; i++ {
		fmt.Print("*")
	}

	fmt.Println("")
}

func emptyLine(width, height int32) {
	var i, j int32

	for j = 0; j < height; j++ {
		fmt.Print("*")

		for i = 0; i < width; i++ {
			fmt.Print(" ")
		}
		fmt.Println("*")
	}
}

func borderLine(width, height int32, leftOrRight bool) {
	var i, j int32

	if leftOrRight {
		for i = 0; i < height; i++ {
			fmt.Println("*")
		}
	} else {
		for i = 0; i < height; i++ {
			for j = 0; j < width; j++ {
				fmt.Print(" ")
			}
			fmt.Println("*")
		}
	}
}

func digitZero(width, height int32) {
	fullLine(width)
	emptyLine(width, height-2)
	fullLine(width)
}
func digitOne(width, height int32) {
	borderLine(width, height, false)
}
func digitTwo(width, height int32) {
	fullLine(width)
	borderLine(width, (height-3)/2, false)
	fullLine(width)
	borderLine(width, (height-3)/2, true)
	fullLine(width)
}
func digitThree(width, height int32) {
	fullLine(width)
	borderLine(width, (height-3)/2, false)
	fullLine(width)
	borderLine(width, (height-3)/2, false)
	fullLine(width)
}

func digitFour(width, height int32) {
	emptyLine(width, (height-1)/2)
	fullLine(width)
	borderLine(width, (height-1)/2, false)
}
func digitFive(width, height int32) {
	fullLine(width)
	borderLine(width, (height-1)/2, true)
	fullLine(width)
	borderLine(width, (height-1)/2, false)
	fullLine(width)
}
func digitSix(width, height int32) {
	fullLine(width)
	borderLine(width, (height-3)/2, true)
	fullLine(width)
	emptyLine(width, (height-3)/2)
	fullLine(width)
}
func digitSeven(width, height int32) {
	fullLine(width)
	borderLine(width, (height - 1), false)
}
func digitEight(width, height int32) {
	fullLine(width)
	emptyLine(width, (height-3)/2)
	fullLine(width)
	emptyLine(width, (height-3)/2)
	fullLine(width)
}
func digitNine(width, height int32) {
	fullLine(width)
	emptyLine(width, (height-3)/2)
	fullLine(width)
	borderLine(width, (height-3)/2, false)
	fullLine(width)
}

func printDigit(x int32) {
	if x == 0 {
		digitZero(10, 15)
	}
	if x == 1 {
		digitOne(10, 15)
	}
	if x == 2 {
		digitTwo(10, 15)
	}
	if x == 3 {
		digitThree(10, 15)
	}
	if x == 4 {
		digitFour(10, 15)
	}
	if x == 5 {
		digitFive(10, 15)
	}
	if x == 6 {
		digitSix(10, 15)
	}
	if x == 7 {
		digitSeven(10, 15)
	}
	if x == 8 {
		digitEight(10, 15)
	}
	if x == 9 {
		digitNine(10, 15)
	}
}

func printNumber(number int32) {

	for {
		if number < 10 {
			printDigit(number)
			break
		}
		printDigit(number % 10)
		number /= 10
	}

}
