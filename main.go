package main

import (
	"clasedh/practica"
	"fmt"
)

func main() {
	slc := []int{2, 3, 5, 1, 6, 3}
	fmt.Printf("1. Suma Slice: %d\n**\n", practica.SumaSlice(slc))

	sl2 := []int{1, 2, 3}

	slcUnido := practica.UnionSlices(slc, sl2)
	fmt.Printf("2. Union de Slices: %v\n**\n", slcUnido)

	slcStr := []string{"rojo", "pelota", "rio", "banana"}
	fmt.Printf("3. Filtrar Slice: %v\n**\n", practica.FiltrarSlice(slcStr, 5))

	slcStr2 := []string{"botella", "roca", "color", "pelota"}
	practica.OrdenarSlice(slcStr2)
	fmt.Printf("4. Ordenar Slice: %v\n", slcStr2)
}
