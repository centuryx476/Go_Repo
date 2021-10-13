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
	// Read input from user
	reader := bufio.NewReader(os.Stdin)

	whatToSay := doctor.Intro()

	fmt.Println(whatToSay)
}

// End of Line
