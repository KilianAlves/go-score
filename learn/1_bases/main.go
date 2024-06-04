// déclaration du package pour ce fichier
package main

// importation du package fmt
import "fmt"

// point d'entrée d'un programme Go
func main() {
	// invocation de la fonction Println déclarée dans le package fmt
	fmt.Println("Hello, World!")
	fizzBuzz(15)
	fmt.Println(isPrime(4))
}

func fizzBuzz(number int) {
	for i := 1; i < number; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func isPrime(number int) bool {
	for i := 1; i < number; i++ {
		for j := 1; j < number; j++ {
			if i*j == number {
				return false
			}
		}
	}
	return true
}
