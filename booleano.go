// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	fmt.Println(ExistWord())
}

func ExistWord() bool {
	frase := "casa"
	palabra := "casa mia"
	contador := 0

	fmt.Println("valida mi casa verde dentro de mi casa verde esta alli", len(palabra))
	fmt.Println("Retornando un valor BOOLEANO")
	fmt.Println("*******")

	for i := 0; i < len(palabra); i++ {

		for j := 0; j < len(frase); j++ {
			//fmt.Println("recorrido dentro del for j", frase[j])
			if palabra[i] == frase[j] {
				contador++
			}

		}

		if contador != len(palabra) {
			return true
		}

	}
	return false
}
