package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func countErrors(files []string) int {
	var wg sync.WaitGroup
	errorChannel := make(chan int, len(files))

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			count := processFile(file)
			errorChannel <- count
		}(file)
	}

	wg.Wait()
	close(errorChannel)

	/*
			"log 1", "log 2", "log 3"

		5-0-3

	*/

	totalErrors := 0
	for count := range errorChannel {
		totalErrors += count
	}

	return totalErrors
}

/*
func analyzeErr (path string){
//corroborar si hay error asignar al booleano true un retorno igual a uno
si no retorno igual a 0
}

boolean X ---> True    0x33-5u3
Y ---> False			00x3F-U7

1 true
0 false

boolean se usan para la comparacion de datos o referencias esperadas a recibir
*/
func processFile(file string) int {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file", file, err)
		return 0
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	errorCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			errorCount++
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", file, err)
	}

	return errorCount

}

func main() {
	files := []string{"C:\\Users\\akali\\GolandProjects\\PracticasCohorte8\\NivelacionConcurrencia\\log1.txt", "log2.txt"}
	totalErrors := countErrors(files)
	fmt.Println("Total Errors: ", totalErrors)
}
