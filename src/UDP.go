package main

import (
	"log"
	"net"
	"sync"
)

func StartUDPProxy(Local, Remote string, WaitGroup *sync.WaitGroup) {

	defer WaitGroup.Done()

	LocalAddr, _ := net.ResolveUDPAddr("udp", Local)
	RemoteAddr, _ := net.ResolveUDPAddr("udp", Remote)

	Conn, Err := net.ListenUDP("udp", LocalAddr)

	if Err != nil {
		log.Printf("[UDP] Erro ao abrir porta %s: %v", Local, Err)
		return
	}

	defer Conn.Close()

	log.Printf("[UDP] Encaminhando: %s -> %s", Local, Remote)

	Buffer := make([]byte, 4096)

	for {

		N, ClientAddr, Err := Conn.ReadFromUDP(Buffer)

		if Err != nil {
			continue
		}

		// Encaminha o pacote em uma nova goroutine
		go func(Payload []byte, Addr *net.UDPAddr) {

			ProxyConn, Err := net.DialUDP("udp", nil, RemoteAddr)

			if Err != nil {
				return
			}

			defer ProxyConn.Close()

			ProxyConn.Write(Payload)

			// Resposta do servidor
			Response := make([]byte, 4096)
			RN, _, Err := ProxyConn.ReadFromUDP(Response)

			if Err == nil {
				Conn.WriteToUDP(Response[:RN], Addr)
			}

		}(append([]byte{}, Buffer[:N]...), ClientAddr)
	}
}
