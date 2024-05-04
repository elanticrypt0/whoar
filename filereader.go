package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type FileReader struct {
	WhoAr               *WhoAr
	filepath            string
	currentDomain       string
	outputPath          string
	bufferSize          int
	domainsInfoFile     string
	domainsInfoFilePath string
}

const defaultBufferSize = 50 * 1024 * 1024

func NewFileReader(war *WhoAr, filepath string) *FileReader {
	return &FileReader{
		WhoAr:               war,
		currentDomain:       "",
		filepath:            filepath,
		outputPath:          "./",
		bufferSize:          defaultBufferSize,
		domainsInfoFile:     "__all_positive_domains_results.txt",
		domainsInfoFilePath: "",
	}
}

func (fr *FileReader) setCurrentDomain(domain string) {
	if fr.WhoAr.IsArDomain(domain) {
		domain = strings.ToLower(domain)
		fr.currentDomain = domain
		fr.WhoAr.SetDomain(domain)
	} else {
		fr.currentDomain = ""
	}
}

func (fr *FileReader) SetOutputPath(path string) {
	if path != "" {
		fr.outputPath = path
		fr.WhoAr.SetOutputPath(path)
	} else {
		fr.outputPath = "./"
	}

	fr.domainsInfoFilePath = fr.outputPath + "/" + fr.domainsInfoFile
}

func (fr *FileReader) SetBufferSize(bufferSize int) {
	fr.bufferSize = bufferSize * 1024 * 1024
}

func (fr *FileReader) SetFilePath(filepath string) {
	fr.filepath = filepath
}

func (fr *FileReader) Run() {

	// Abrir el archivo de entrada
	file, err := os.Open(fr.filepath)
	if err != nil {
		fmt.Printf("Error al abrir el archivo: %v\n", err)
		return
	}
	defer file.Close()

	// Crear slice para almacenar todos los dominios
	allDomains := []string{}

	// Crear scanner para leer el archivo línea por línea
	scanner := bufio.NewScanner(file)

	// Configurar el tamaño máximo del buffer del scanner
	scanner.Buffer(make([]byte, fr.bufferSize), fr.bufferSize)

	// Leer el archivo línea por línea
	for scanner.Scan() {
		line := scanner.Text()

		line_trimmed := strings.TrimSpace(line)

		fr.setCurrentDomain(line_trimmed)

		allDomains = append(allDomains, fr.currentDomain)

		if fr.currentDomain != "" {
			fr.WhoAr.Run()
		}
	}

	// Verificar errores de escaneo
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error al leer el archivo línea por línea: %v\n", err)
		return
	}

	fr.SaveAllDomainsInfo(allDomains)

}

func (fr *FileReader) SaveAllDomainsInfo(allDomainsInfo []string) {
	// Escribir los dominios inválidos en el archivo CSV
	allDomainsCSVFile, err := os.Create(fr.domainsInfoFilePath)

	if err != nil {
		fmt.Printf("Error al crear el archivo de resultados: %v\n", err)
		return
	}
	defer allDomainsCSVFile.Close()

	allDomainsWriter := csv.NewWriter(allDomainsCSVFile)
	defer allDomainsWriter.Flush()

	for _, domain := range allDomainsInfo {
		if err := allDomainsWriter.Write([]string{domain}); err != nil {
			fmt.Printf("Error al escribir en el archivo CSV: %v\n", err)
			return
		}
	}

	fmt.Println("Lista de resultados positivos: ", fr.domainsInfoFile)
}
