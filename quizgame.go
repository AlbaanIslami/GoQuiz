package main

import "fmt"

func main() {
	fmt.Println("Welcome to the quiz")

	fmt.Printf("Enter you name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcolme to the game\n", name)

	fmt.Printf("enter your age: ")
	var age uint
	fmt.Scan(&age)

	fmt.Println(age >= 20)
	if age >= 20 {
		fmt.Println("You can play")
	} else {
		fmt.Println("You cant play")
		return
	}

	score := 0
	num_questions := 3

	fmt.Printf("What is better, RTX 3080 or RTX 3090?")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3090" || answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("What is better, Mercedes or BMW? ")
	var answer3 string
	fmt.Scan(&answer3)

	if answer3 == "Mercedes" || answer3 == "mercedes" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many wheels are on a standard car? ")
	var answer4 uint
	fmt.Scan(&answer4)

	if answer4 == 4 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v", score, num_questions)
	percent := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v %% ", percent)
}
