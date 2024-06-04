// déclaration du package pour ce fichier
package main

// importation du package fmt
import "fmt"

// point d'entrée d'un programme Go
func main() {
	// invocation de la fonction Println déclarée dans le package fmt
	fmt.Println("Hello, World!")
	fizzBuzz(6)
	fizzBuzz(10)
	fizzBuzz(15)
}

func fizzBuzz(number int) {
	switch {
	case number%3 == 0 && number%5 == 0:
		fmt.Println("FizzBuzz")
	case number%3 == 0:
		fmt.Println("Fizz")
	case number%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(number)
	}
}
