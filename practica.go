package main

import "fmt"
import "strings" 
import "strconv" 

func main() {
	readNmber()
	isPair()
	daysToSeconds()
	average() 
	calcPerimetro()
	secondsToDays()
}



//1. Realizar un algoritmo para leer un número e informar si es mayor, igual o menos a cero.
func readNmber()  {
	var num int

    fmt.Print("Ingrese el numero: ")
    fmt.Scanln(&num)

	if num > 0 {
		fmt.Println("El numero " , num , " es mayor que 0")
	}else if num == 0{
		fmt.Println("El numero " , num , " es mayor que 0")
	}else{
		fmt.Println("El numero " , num , " es menor que 0")
	}
}

//2. Escribir un algoritmo que determine si un número es par.

func isPair()  {

	var num int

    fmt.Print("Ingrese el numero: ")
    fmt.Scanln(&num)

	if num % 2 == 0 {
		fmt.Println("El numero es par")
	}else {
		fmt.Println("El numero es impar")
	}
}
//3. Dados tres valores ambientales de temperatura, desarrollar un algoritmo para calcular e informar la suma y el promedio de dichos valores.

func average()  {
	var value,value2,value3 float64

    fmt.Print("Ingrese el primer valor de temperatura: ")
    fmt.Scanln(&value)

	fmt.Print("Ingrese el segundo valor de temperatura: ")
    fmt.Scanln(&value2)

	fmt.Print("Ingrese el tecer valor de temperatura: ")
    fmt.Scanln(&value3)

	sum := value + value2 + value3
	average := sum / 3
	fmt.Println("El promedio de temperatura es: ", average)
}


//4. Redactar un algoritmo para pasar un periodo expresado en días, horas, minutos y segundos a periodo expresado en segundos.

func daysToSeconds()  {
	var date string

    fmt.Print("Ingrese el valor en días, horas, minutos y segundos ejemplo del formato 18/23/12/50: ")
    fmt.Scanln(&date)
	var arregloInts []int
	valueSplit := strings.Split(date, "/")
	for _, str := range valueSplit {

		numero, _ := strconv.Atoi(str)
		arregloInts = append(arregloInts, numero)
	}

	seconds := arregloInts[0] * 86400
	seconds = seconds + (arregloInts[1] * 3600)
	seconds = seconds + (arregloInts[2] * 60)
	seconds = seconds + arregloInts[3]

	fmt.Println("El periodo en segundos equivale a: ", seconds)
}


//5. Expresar a un número cualquiera como la suma de una serie de números siguiendo las restricciones impuestas a continuación.z




//6.Para valientes: Pasar un periodo expresado en segundos a un periodo expresado en días, horas, minutos y segundos.

func secondsToDays()  {
	var secondsStr string

    fmt.Print("Ingrese el valor que desea convertir: ")
    fmt.Scanln(&secondsStr)

	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		fmt.Println("Error al convertir el valor a entero:", err)
		return
	}

	var days,hours,minutes int

	days = seconds / 86400
	seconds = seconds % 86400
	if seconds >= 3600{
		hours = seconds / 3600
		seconds = seconds % 3600
	}

	if seconds >= 60{
		minutes = seconds / 60
		seconds = seconds % 60
	}

	fmt.Println("El periodo en Dias equivale a: ", days)
	fmt.Println("El periodo en Horas equivale a: ", hours)
	fmt.Println("El periodo en minutps equivale a: ", minutes)
	fmt.Println("El periodo en segundos equivale a: ", seconds)
}
 


/*Descripción del Problema:
Escribe un programa que calcule el área y el perímetro de un rectángulo. El programa debe pedir al usuario que introduzca la longitud y la anchura del rectángulo y luego calcular y mostrar el área y el perímetro.
Especificaciones:
Solicita al usuario que ingrese la longitud y la anchura del rectángulo.

Calcula el área del rectángulo (longitud * anchura).
Calcula el perímetro del rectángulo (2 * (longitud + anchura)).
Muestra los resultados (área y perímetro) al usuario.
Ejemplo de Salida:
Ingrese la longitud del rectángulo: 5 
Ingrese la anchura del rectángulo: 3 


El área del rectángulo es: 15 
El perímetro del rectángulo es: 16*/

func calcPerimetro()  {
	var longitud,ancho float32

    fmt.Print("Ingrese por favor el valor de la longitud: ")
    fmt.Scanln(&longitud)

	fmt.Print("Ingrese por favor el valor del ancho: ")
    fmt.Scanln(&ancho)


	area := longitud * ancho
	perimetro := (longitud*2) + (ancho*2)

	fmt.Println("Valor del area: ",area)
	fmt.Println("Valor del Perimetro: ", perimetro)

}