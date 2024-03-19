package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/ubtref-ayp2/mymodules/math"
	"github.com/ubtref-ayp2/mymodules/saludos"

)

func main(){
	fmt.Println("Hola mundo!")
	msg := color.CyanString("Hola Go Modules")
	fmt.Println(msg)
	a := 10
	b := 5
	c := math.Sumar(a, b)
	fmt.Printf("La suma de %d y %d es %d\n", a, b, c)
	fmt.Printf("El valor de PI es: %f\n", math.PI)
	fmt.Println(saludos.Saludando())
}

