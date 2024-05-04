package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type WhoAr struct {
	domain         string
	outputPath     string
	output         []byte
	isActiveDomain bool
}

func NewWhoAr() *WhoAr {
	return &WhoAr{
		isActiveDomain: false,
	}
}

func (whoar *WhoAr) SetDomain(domain string) {
	whoar.domain = domain
}

func (whoar *WhoAr) SetOutputPath(path string) {
	whoar.outputPath = path
}

func (whoar *WhoAr) Run() {
	if whoar.IsArDomain(whoar.domain) {
		whoar.whois()
		if whoar.isActiveDomain {
			whoar.saveFile()
		} else {
			log.Println("El dominio: ", whoar.domain, " No se encuentra activo")
		}
	} else {
		log.Println("El dominio: ", whoar.domain, " no termina en .AR")
	}

}

func (whoar *WhoAr) IsArDomain(domain string) bool {
	isAr := false
	if strings.HasSuffix(domain, ".ar") {
		return true
	}
	return isAr
}

func (whoar *WhoAr) whois() {
	// Ejecutar el script bash
	cmd := exec.Command("whois", whoar.domain)
	output, err := cmd.Output()
	if err != nil {
		log.Println("Error al ejecutar el script:", err)
		os.Exit(1)
	}
	// Comparar la cadena
	if string(output) != "El dominio no se encuentra registrado en NIC Argentina\n" {
		whoar.isActiveDomain = true
		whoar.output = output
	}
}

func (whoar *WhoAr) saveFile() {
	// Guardar la salida en un archivo con el nombre del dominio + .txt
	dominioFile := strings.ReplaceAll(whoar.domain, ".", "_")
	nombreArchivo := dominioFile + ".txt"
	if whoar.outputPath != "" {
		nombreArchivo = whoar.outputPath + "/" + nombreArchivo
	}
	archivo, err := os.Create(nombreArchivo)
	if err != nil {
		log.Println("Error al crear el archivo:", err)
		os.Exit(1)
	}
	defer archivo.Close()

	_, err = archivo.Write(whoar.output)
	if err != nil {
		log.Println("Error al escribir en el archivo:", err)
		os.Exit(1)
	}

	log.Println("Salida del script guardada en", nombreArchivo)
}
