package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Uso: portproxy.exe [HostDestino] [Protocolo]/[PortaLocal]:[PortaRemota]")
		fmt.Println("Ex:  portproxy.exe 192.168.1.10 tcp/80:80 udp/53:53")
		return
	}

	LogFile := SetupLogging()
	defer LogFile.Close()

	TargetHost := os.Args[1]
	Mappings := os.Args[2:]

	var WaitGroup sync.WaitGroup

	for _, Entry := range Mappings {

		// Formato esperado: protocolo/portaLocal:portaRemota
		Parts := strings.Split(Entry, "/")

		if len(Parts) != 2 {
			log.Printf("Formato inválido: %s\n", Entry)
			continue
		}

		Protocol := strings.ToLower(Parts[0])
		Ports := strings.Split(Parts[1], ":")

		if len(Ports) != 2 {
			log.Printf("Portas inválidas em: %s\n", Entry)
			continue
		}

		LocalAddr := ":" + Ports[0]
		RemoteAddr := TargetHost + ":" + Ports[1]

		WaitGroup.Add(1)

		switch Protocol {
		case "tcp":
			go StartTCPProxy(LocalAddr, RemoteAddr, &WaitGroup)

		case "udp":
			go StartUDPProxy(LocalAddr, RemoteAddr, &WaitGroup)
		}
	}

	log.Printf("Proxy iniciado para o Host: %s\n", TargetHost)

	WaitGroup.Wait()
}
