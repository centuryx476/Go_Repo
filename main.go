package main // This can be named whatever you want

import "fmt" // Import "library" for output

// Func has to be named 'main'
func main() {
	fmt.Println("Hello from Go!") // Prints on NewLine

	// Call the below function
	sayHelloWorld("Passing string to function for output")
}

// Function for Output
func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}

// End of Line
