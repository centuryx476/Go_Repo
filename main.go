package main // This can be named whatever you want

// Importing multiple
import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

// Func has to be named 'main'
func main() {
	// Variable to read input from user
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	// In Go there is only one type of Loop (To get multiple inputs)
	for {
		// Adding a little prompt for the screen
		fmt.Print("-> ")

		// Actually reading input from user (Only one word it can read with \n)
		userInput, _ := reader.ReadString('\n')

		// This helps make the break from the loop
		userInput = strings.Replace(userInput, "\r\n", "", -1) // Replace return for Windows
		userInput = strings.Replace(userInput, "\n", "", -1)   // Replace return for everything else

		// TODO: I want to research a way to do away with the egyptian style formatting.
		// I find it really annoying

		// Checking for 'quit' so we can exit loop
		if userInput == "quit" {
			break
		} else {
			// Output the response from eliza
			fmt.Println(doctor.Response(userInput))
		}
	}
}

// This code will be removed and starting fresh with the new section

//End of Line
