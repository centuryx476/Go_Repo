package main // This can be named whatever you want

import (
	"fmt"
	"myapp/doctor"
)

// Func has to be named 'main'
func main() {
	var whatToSay string

	whatToSay = doctor.Intro()
	fmt.Println(whatToSay)

}
