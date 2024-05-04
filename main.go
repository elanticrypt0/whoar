package main

import (
	"flag"
	"fmt"
)

func main() {

	AppBanner()

	inputFile := flag.String("f", "", "Ingresa una lista de dominios")
	domainAr := flag.String("d", "", "Debes ingresar un dominio .ar")
	outputPath := flag.String("o", "", "Ingresa un directorio de salida")
	bufferSize := flag.Int("buffer", 50, "Setea el tamaño del buffer dependiendo del archivo. Por defecto es de 50 MB")

	flag.Parse()

	if *domainAr == "" && *inputFile == "" {
		fmt.Println("Debes ingresar un dominio utilizando -d [DOMINIO_AR]")
		fmt.Println("O una lista de dominios con -f [ARCHIVO]")
		return
	}

	dominio := *domainAr
	salida := *outputPath

	whoar := NewWhoAr()

	if *outputPath != "" {
		whoar.SetOutputPath(salida)
	}

	if *inputFile != "" {
		filereader := NewFileReader(whoar, *inputFile)
		filereader.SetOutputPath(salida)
		if *bufferSize != 50 {
			filereader.SetBufferSize(*bufferSize)
		}
		filereader.Run()
	} else {
		whoar.SetDomain(dominio)
		if *outputPath != "" {
			whoar.SetOutputPath(salida)
		}
		whoar.Run()
	}

}
