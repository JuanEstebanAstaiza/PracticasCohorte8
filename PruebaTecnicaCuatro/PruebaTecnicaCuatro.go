package main

/*import (
	"fmt"
	"math/rand"
	"time"
)

func generarMatriz(filas, columnas int) [][]int {
	matriz := make([][]int, filas)
	for i := 0; i < filas; i++ {
		matriz[i] = make([]int, columnas)
		for j := 0; j < columnas; j++ {
			matriz[i][j] = rand.Intn(10)
		}
	}
	return matriz
}

// Funcion de imprimir la matriz

func imprimirMatriz(matriz [][]int) {

	for _, fila := range matriz {
		fmt.Println(fila)
	}
}

//Funcion de multiplicacion

func multiplicarMatricesConcurrente(matrizA, matrizB [][]int, resultado chan<- [][]int) {
	filasA := len(matrizA)
	columnasA := len(matrizA[0])
	columnasB := len(matrizB[0])

	resultadoMatriz := make([][]int, filasA)
	for i := 0; i < filasA; i++ {
		resultadoMatriz[i] = make([]int, columnasB)
	}

	for i := 0; i < filasA; i++ {
		for j := 0; j < columnasB; j++ {
			go func(i, j int) {
				suma := 0
				for k := 0; k < columnasA; k++ {
					suma += matrizA[i][k] * matrizB[k][j]
				}
				resultadoMatriz[i][j] = suma
			}(i, j)
		}
	}
	for x := 0; x < filasA; x++ {
		for z := 0; z < columnasB; z++ {
			<-time.After(10 * time.Millisecond)
		}
	}

	resultado <- resultadoMatriz
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//definir tamaño de matriz

	filasA := 3
	columnasA := 2
	filasB := 2
	columnasB := 3

	matrizA := generarMatriz(filasA, columnasA)
	matrizB := generarMatriz(filasB, columnasB)

	//Upwork
	fmt.Println("Matriz A: ")
	imprimirMatriz(matrizA)
	fmt.Println("Matriz B: ")
	imprimirMatriz(matrizB)

	resultado := make(chan [][]int)

	go multiplicarMatricesConcurrente(matrizA, matrizB, resultado)

	fmt.Println("Resultado del producto cruz: ")
	fmt.Println(<-resultado)
}*/

//indeed

/*import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Competidor struct {
	nombre    string
	velocidad int
}

func correrCompetidor(competidor Competidor, pista chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for distancia := 0; distancia < 100; {
		avance := rand.Intn(competidor.velocidad)
		pista <- avance
		distancia += avance
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("%s ha terminas la carrera! \n", competidor.nombre)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	competidores := []Competidor{
		{"Competidor 1", rand.Intn(10) + 1},
		{"Competidor 2", rand.Intn(10) + 1}, // Velocidades entre el 1 y 10
		{"Competidor 3", rand.Intn(10) + 1},
		{"Competidor 4", rand.Intn(10) + 1},
	}

	pista := make(chan int)

	defer close(pista)

	var wg sync.WaitGroup
	wg.Add(len(competidores))

	for _, competidor := range competidores {
		go correrCompetidor(competidor, pista, &wg)
	}

	for {
		select {
		case avance := <-pista:
			fmt.Printf("Avance: %d\n", avance)
		case <-time.After(1 * time.Second):
			fmt.Println("No hay avance en 1 segundo...")

		}

		select {
		case <-pista:
			fmt.Println("!Tenemos un ganador!")
			wg.Wait()
			return
		default:
		}
	}
}*/

/*import (
	"fmt"
	"time"
)

// funcion para simular tareas

func tarea(id int, canal chan string) {

	//simular trabajo
	time.Sleep(time.Second * time.Duration(id))

	//enviar mensaje

	canal <- fmt.Sprintf("tarea %d completada", id)
}

func main() {
	// crear el canal
	canal := make(chan string)

	//definir el numero de tareas
	numTareas := 5

	//inicar goroutines separadas
	for i := 1; i <= numTareas; i++ {
		go tarea(i, canal)
	}

	//esperar a que terminen las go routines
	for x := 1; x < numTareas; x++ {
		fmt.Println(<-canal)
	}
}*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Función para producir fabricar numeros aleatorios y enviarlos en un canal

func productor(numeros chan<- int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 10; i++ {
		numero := rand.Intn(100)
		numeros <- numero
		fmt.Printf("Productor: Enviado número %d\n", numero)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond) //Simulamos una produccion irregular
	}

	close(numeros)
}

//Funcion  consumidora que recibe numeros del canal y los procesa

func consumidor(numeros <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for numero := range numeros {

		fmt.Printf("Consumidor: Recibido número %d\n", numero)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func consumidorSuma(numeros <-chan int, wg *sync.WaitGroup) {

}

func main() {
	rand.Seed(time.Now().UnixNano())

	//Creamos el canal de comunicacion entre productores y consumidores

	numeros := make(chan int)

	// Waitgropu para esperar la finalizacion de las goroutines
	var wg sync.WaitGroup

	// Iniciamos las goroutines de productor y consumidor

	wg.Add(3)
	go productor(numeros, &wg)
	go consumidor(numeros, &wg)
	go consumidorSuma(numeros, &wg)
	//Esperamos a que terminen

	wg.Wait()
}
