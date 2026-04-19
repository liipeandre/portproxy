package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func SetupLogging() *os.File {

	ExePath, Err := os.Executable()

	if Err != nil {
		log.Printf("Erro ao obter o caminho do executável: %v\n", Err)
		os.Exit(1)
	}

	ExeName := filepath.Base(ExePath)
	NameWithoutExt := strings.TrimSuffix(ExeName, filepath.Ext(ExeName))

	TargetDir := `C:\Logs\portproxy`

	LogPath := filepath.Join(TargetDir, NameWithoutExt+".log")

	Err = os.MkdirAll(TargetDir, 0755)

	if Err != nil {
		log.Printf("Erro ao criar o diretório de log em %s: %v\n", TargetDir, Err)
		os.Exit(1)
	}

	File, Err := os.OpenFile(LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if Err != nil {
		log.Printf("Erro ao abrir arquivo de log em %s: %v\n", LogPath, Err)
		os.Exit(1)
	}

	log.SetOutput(File)

	return File
}
