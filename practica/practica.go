package practica

// 1 Suma de los elementos
// Dado un slices de numeros enteros, devolver la suma de todos los elementos
// Entrada: [2, 3, 5, 1, 6, 3]
// Salida: 20

func SumaSlice(slice []int) int {
	suma := 0
	for i := range slice {
		suma += slice[i]
	}
	return suma
}

// 2 Unir dos Slices
// Dados dos slices, crear un nuevo slice que contenga todos los elementos de ambos slices, sin duplicados.
// Entrada: [1, 2, 3] , [4, 5, 6]
// Salida: [1, 2, 3, 4, 5, 6]

func UnionSlices(slice1 []int, slice2 []int) []int {
	resultado := slice1
	resultado = append(resultado, slice2...)
	return resultado
	// return append(slice1, slice2...)
}

// 3 Filtrar por longitud
// Dado un slice de strings con palabras y un numero de longitud minima, devolver un slice con las palabras que tengan una longitud mayor al numero dado.
// Slice Entrada: ["rojo", "pelota", "rio", "banana"]
// Longitud minima: 5
// Salida: ["pelota", "banana"]

func FiltrarSlice(slice []string, min int) []string {
	resultado := make([]string, 0, len(slice))
	for _, v := range slice {
		if len(v) >= min {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

// 4 Ordenar palabras por longitud
// Dado un slice de strings, ordenar las palabras por longitud en orden ascendente (de menor a mayor)
// Slice Entrada: ["botella", "roca", "color", "pelota" ]
// Salida: ["roca", "color", "pelota", "botella"]

func OrdenarSlice(slice []string) {
	aux := ""
	for i := range slice {
		for j := range slice {
			if len(slice[i]) < len(slice[j]) {
				aux = slice[i]
				slice[i] = slice[j]
				slice[j] = aux
			}
		}
	}

}

// Extra - Eliminar duplicados
// Dado un slice de numeros enteros, devolver un nuevo slice sin duplicados.
// Entrada: [2, 3, 6, 2, 4, 6, 0, 3]
// Salida: [2, 3, 6, 4, 0]

func EliminarDuplicadosSlice(slice []int) []int {
	unicos := make(map[int]bool)
	resultado := make([]int, 0, len(slice))
	for _, numero := range slice {
		if !unicos[numero] {
			unicos[numero] = true
			resultado = append(resultado, numero)
		}
	}
	return resultado
}
