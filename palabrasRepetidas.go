package main

import "fmt"

func LetrasRepetidas() {
	array1 := []string{"d", "d", "d", "d", "b", "b", "b", "c", "c", "c", "g", "g", "e", "d"}
	var array2 [][]string
	for i := 0; i < len(array1); i++ {
		var arrayTemporal []string
		for j := 0; j < len(array1); j++ {
			if array1[i] == array1[j] {
				arrayTemporal = append(arrayTemporal, array1[j])
			}
		}
		if len(arrayTemporal) > 1 {
			if len(array2) == 0 {
				array2 = append(array2, arrayTemporal)
			} else {
				if validarLetrasRepetidas(array2, arrayTemporal) == true {
					array2 = append(array2, arrayTemporal)
				}
			}
		}
	}
	fmt.Println(array2)
}

func validarLetrasRepetidas(array1 [][]string, array2 []string) bool {
	cont := 0
	var valido bool
	valido = true
	for i := 0; i < len(array1); i++ {
		if array1[i][0] == array2[0] {
			cont++
		}
	}
	if cont > 0 {
		valido = false
	}
	return valido
}

// Programa principal
func main() {
	LetrasRepetidas()
}
