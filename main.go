// main.go
package main

import "fmt"

// Greet returns a greeting message for v2
func Greet(name string) string {
	return fmt.Sprintf("Greetings, %s! (from v2)", name) // Changed message
}

func main() {
	fmt.Println(Greet("World"))
}
