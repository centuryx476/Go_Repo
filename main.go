package main // This can be named whatever you want

// Importing multiple
import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
)

// Func has to be named 'main'
func main() {
	// Variable to read input from user
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()
	fmt.Println(whatToSay)

	// In Go there is only one type of Loop (To get multiple inputs)
	for {
		// Actually reading input from user (Only one word it can read with \n)
		userInput, _ := reader.ReadString('\n')

		// Output the userInput
		fmt.Println(userInput)
	}
}

// End of Line
