// déclaration du package pour ce fichier
package main

// importation du package fmt
import (
	"fmt"
)

// point d'entrée d'un programme Go
func main() {
	// invocation de la fonction Println déclarée dans le package fmt
	fmt.Println("Hello, World!")
	fizzBuzz(15)
	fmt.Println(isPrime(1))
	fmt.Println(primeNumbers(50))
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
	if number != 2 && number%2 == 0 || number < 2 {
		return false
	}

	for i := 3; i < number; i += 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func primeNumbers(number int) []int {
	var primeNumbers []int
	for i := 1; i <= number; i++ {
		if isPrime(i) {
			primeNumbers = append(primeNumbers, i)
		}
	}
	return primeNumbers
}
