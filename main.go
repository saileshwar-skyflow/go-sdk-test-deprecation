// main.go
package main

import "fmt"

// Greet returns a greeting message for v1
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s! (from v1)", name)
}

func main() {
	fmt.Println(Greet("World"))
}
