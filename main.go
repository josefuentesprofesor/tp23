package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Abre el archivo CSV
	file, err := os.Open("GOOG.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crea un lector CSV
	reader := csv.NewReader(file)

	// Lee todas las filas del archivo
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}

	// Inicializa variables para el cálculo
	totalVolume := 0
	maxVolume := 0
	minVolume := -1 // Inicializado con un valor negativo para asegurarnos de que cualquier valor sea menor

	// Itera sobre las filas y calcula el promedio, máximo y mínimo de la columna "Volume"
	for _, row := range rows[1:] { // Empezamos desde la segunda fila para omitir la fila de encabezados
		volume, err := strconv.Atoi(row[6])
		if err != nil {
			fmt.Println("Error al convertir el volumen a entero:", err)
			return
		}

		//TODO actualizar el acumulador

		//TODO actualizar el maximo

		//TODO actualizar el minimo
	}

	//TODO Calcular el promedio

	// Imprime los resultados
	fmt.Printf("Promedio de Volume: %.2f\n", averageVolume)
	fmt.Printf("Máximo de Volume: %d\n", maxVolume)
	fmt.Printf("Mínimo de Volume: %d\n", minVolume)
}
