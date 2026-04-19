package main

import (
	"io"
	"log"
	"net"
	"sync"
)

func StartTCPProxy(Local, Remote string, WaitGroup *sync.WaitGroup) {

	defer WaitGroup.Done()

	Listener, Err := net.Listen("tcp", Local)

	if Err != nil {
		log.Printf("[TCP] Erro ao abrir porta %s: %v", Local, Err)
		return
	}

	log.Printf("[TCP] Encaminhando: %s -> %s", Local, Remote)

	for {

		Conn, Err := Listener.Accept()

		if Err != nil {
			log.Printf("[TCP] Erro na conexão: %v", Err)
			continue
		}

		go HandleTCPStream(Conn, Remote)
	}
}

func HandleTCPStream(LocalConn net.Conn, RemoteAddr string) {

	defer LocalConn.Close()

	RemoteConn, Err := net.Dial("tcp", RemoteAddr)

	if Err != nil {
		log.Printf("[TCP] Erro ao conectar no destino %s: %v", RemoteAddr, Err)
		return
	}

	defer RemoteConn.Close()

	// Ponte bidirecional
	Done := make(chan bool, 2)

	Copy := func(Dst, Src net.Conn) {
		io.Copy(Dst, Src)
		Done <- true
	}

	go Copy(RemoteConn, LocalConn)
	go Copy(LocalConn, RemoteConn)

	<-Done
}
