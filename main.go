package main

import (
	"fmt"
	"sort"
	"os"
)

func archivo(cadena []string){
	file1, err1 := os.Create("ascendente.txt")
	if err1 !=nil{
		fmt.Printf("Daño en archivo / no se encuentra el archivo")
		return
	}

	defer file1.Close()

	sort.Strings(cadena)
	for _, j:= range cadena{
		file1.WriteString(j)
	}

	file2, err2 := os.Create("descendente.txt")
	if err2 !=nil{
		fmt.Printf("Daño en archivo / no se encuentra el archivo ")
		return
	}

	defer file2.Close()

	sort.Sort(sort.Reverse(sort.StringSlice(cadena)))
	for _, j:= range cadena{
		file2.WriteString(j)
	}
}

func main(){
	var cadena string
	var n int
	sliceCadena := []string{}
	
	fmt.Printf("Ingrese el numero de cadenas: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		
		fmt.Scan(&cadena)
		sliceCadena = append(sliceCadena, cadena + "\n")
	}
	archivo(sliceCadena)
}