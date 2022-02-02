package main

import "fmt"


func main()  {

	fmt.Println(uno(5))

	//Crear e inicializar dos variables a partir del resultado de la función dos()
	numero, estado := dos(1)

	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println(Calculo(5, 46))
}	

//Función que recibe un entero y devuelve un entero
func uno(numero int) int {
	return numero*2
}

//Función que recibe un entero y devuelve un entero, podemos inicializar variables en la definición de parámetros de salida
func uno1(numero int) (resultado int) {
	resultado = numero*2
	return resultado
}

//Función que recibe un entero y devuelve dos tipos de datos, un entero y un booleano
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	}
	return 10, true
}

func Calculo(numero ......int) int {
	total:=0
	//range es una instrucción que siempre devuelve dos valores, primero devuelve el número de elemento
	//como no vamos a usar el primer valor que devuelve la instrucción range, usamos "_"
	for _, num:= range numero {
		total = total + num
	}
	return total
}