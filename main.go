package main // This can be named whatever you want

import "fmt" // Import "library" for output

// Func has to be named 'main'
func main() {
	fmt.Println("Hello from Go!") // Prints on NewLine

	// One way of creating a variable
	var whatToSay string
	whatToSay = "Passing string to function for output"

	// Call the below function
	sayHelloWorld(whatToSay)
}

// Function for Output
func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}
