package main

import (
	"fmt"
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
	domainFile     string
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
		whoar.makeDomainFile()

		// si el archivo no existe entonces ejecuta todo lo demás.
		if whoar.domainFileNoExists() {
			whoar.whois()
			if whoar.isActiveDomain {
				whoar.saveFile()
			} else {
				log.Println("El dominio: ", whoar.domain, " No se encuentra activo")
			}
		} else {
			log.Panicln("Dominio: ", whoar.domain, " ya fue escaneado.")
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

	msgExceededQueries := "Excediste la cantidad permitida de consultas. Volvé a intentarlo más tarde\n"

	if string(output) != msgExceededQueries {
		// Comparar la cadena
		if string(output) != "El dominio no se encuentra registrado en NIC Argentina\n" {
			whoar.isActiveDomain = true
			whoar.output = output
		}
	} else {
		fmt.Println("")
		fmt.Println(msgExceededQueries)
		log.Panicln("")
	}
}

func (whoar *WhoAr) saveFile() {
	archivo, err := os.Create(whoar.domainFile)
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

	log.Println("Salida del script guardada en", whoar.domainFile)
}

func (whoar *WhoAr) makeDomainFile() {
	// Guardar la salida en un archivo con el nombre del dominio + .txt
	dominioFile := strings.ReplaceAll(whoar.domain, ".", "_")
	nombreArchivo := dominioFile + ".txt"
	if whoar.outputPath != "" {
		nombreArchivo = whoar.outputPath + "/" + nombreArchivo
	}
	whoar.domainFile = nombreArchivo
}

func (whoar *WhoAr) domainFileNoExists() bool {
	if _, err := os.Stat(whoar.domainFile); os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}
