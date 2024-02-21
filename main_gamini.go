// Code generated by Google Gemini (https://gemini.google.com/)
// Some edits were applied to make the code work.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	quitCommand      = "quit"
	coffeeSimulation = "coffee"
	capitalsQuiz     = "capitals"
)

var countryCapitals = map[string]string{
	"France":         "Paris",
	"Germany":        "Berlin",
	"Italy":          "Rome",
	"Spain":          "Madrid",
	"United Kingdom": "London",
	"India":          "New Delhi",
	"Japan":          "Tokyo",
	"Brazil":         "Brasilia",
}

func greetUser() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What's your name? ")
	scanned := scanner.Scan()
	if !scanned {
		fmt.Println("Error reading input:", scanner.Err())
		return ""
	}
	name := strings.TrimSpace(scanner.Text())
	fmt.Printf("Hello, %s!\n", name)
	return name
}

func getTodayInfo() (string, string, string) {
	now := time.Now()
	dayOfWeek := now.Weekday().String()
	date := now.Format("2006-01-02")
	currTime := now.Format("15:04:05")
	return dayOfWeek, date, currTime
}

func brewCoffee(strength int) {
	switch strength {
	case 1:
		fmt.Println("Brewing weak coffee...")
		time.Sleep(time.Second)
		fmt.Println("Your weak coffee is ready!")
	case 2:
		fmt.Println("Brewing medium coffee...")
		time.Sleep(2 * time.Second)
		fmt.Println("Your medium coffee is ready!")
	case 3:
		fmt.Println("Brewing strong coffee...")
		time.Sleep(3 * time.Second)
		fmt.Println("Your strong coffee is ready! Be careful, it's hot!")
	default:
		fmt.Println("Invalid coffee strength. Please choose 1, 2, or 3.")
	}
}

func capitalsQuizF() {
	score := 0
	questionsAsked := 0

	for country, capital := range countryCapitals {
		questionsAsked++
		fmt.Printf("What is the capital of %s? ", country)
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		answer = strings.TrimSpace(answer)

		if strings.EqualFold(answer, capital) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Incorrect. The capital of %s is %s.\n", country, capital)
		}
	}

	fmt.Printf("You answered %d out of %d questions correctly.\n", score, questionsAsked)
}

func displayCommands() {
	fmt.Println(" Available commands:")
	fmt.Println("--------------------")
	fmt.Println("- quit: Close the application.")
	fmt.Println("- coffee: Simulate brewing coffee (weak, medium, or strong).")
	fmt.Println("- capitals: Test your knowledge of world capitals with a quiz.")
	fmt.Println("- date: Display the current date, time, and day of the week.")
	fmt.Println("") // Add some blank lines for visual separation
}

func main() {
	name := greetUser()

	displayCommands()

	for {
		fmt.Print("Enter a command (quit, coffee, capitals, date): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanned := scanner.Scan()
		if !scanned {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}
		command := strings.ToLower(scanner.Text())

		switch command {
		case quitCommand:
			fmt.Println("Goodbye, ", name, "!")
			return

		case coffeeSimulation:
			fmt.Print("Enter coffee strength (1 = weak, 2 = medium, 3 = strong): ")
			scanned := scanner.Scan()
			if !scanned {
				fmt.Println("Error reading input:", scanner.Err())
				continue
			}
			strength, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			brewCoffee(strength)

		case capitalsQuiz:
			capitalsQuizF()

		case "date":
			var dayOfWeek, date, currTime = getTodayInfo() // Now the function is called!
			fmt.Printf("Today is %s, %s, %s\n", dayOfWeek, date, currTime)

		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}
