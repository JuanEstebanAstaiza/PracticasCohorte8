package PracticaUno

/*import (
	"fmt"
	"strings"
)

type Player struct {
	Name       string
	Identifier int
}

func main() {

	//Lista de jugadores actuales
	players := []string{"John Smith", "David Johnson", "Michael Garcia", "Sarah Williams", "Daniel Martinez", "Emily Brown", "James Rodriguez", "Sophia Lee", "Lucas Oliveira", "Olivia Taylor", "Mateo Hernandez", "Emma Wilson", "Gabriel Silva"}

	//registro de nuevos jugadores
	newPlayer := []string{"New Player 1", "New Player 2", "New Player 3"}

	//ELiminar los 3 primeros jugadores
	players = players[3:]

	//agregar
	for _, name := range newPlayer {
		players = append(players, name)
	}

	//mostrar la lista
	fmt.Println("Lista completa: ")
	for i, player := range players {
		fmt.Printf("%d. %s\n", i+1001, sanitizeName(player))
	}
}

//funcionEliminarNombre

func sanitizeName(name string) string {
	//eliminar caracteres especiales
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == ' ' {
			return r
		}
		return -1

	}, name)
}*/

/*import "fmt"

func main() {

	//definir el valor de N

	n := 50

	//imprimir los numeros fizzbuzz

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

		if i < n {
			fmt.Print(", ")
		}

	}
}*/
