package main

import (
	"flag"
	"io"
	"log"
	"net"
)

func handleConnection(clientConn net.Conn, targetAddress string) {
	defer clientConn.Close()

	targetConn, err := net.Dial("tcp", targetAddress)
	if err != nil {
		log.Printf("Failed to connect to target %s: %v", targetAddress, err)
		return
	}
	defer targetConn.Close()

	go func() {
		// Forward data from client to target
		if _, err := io.Copy(targetConn, clientConn); err != nil {
			//if err is connection reset by peer, we can ignore it
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Printf("Connection reset by peer: %v", err)
				return
			}

			//if err is use of closed network connection, we can ignore it
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				log.Printf("Temporary network error: %v", err)
				return
			}
			// Log other errors
			log.Printf("Error forwarding data from client to target: %v", err)
		}
	}()

	// Forward data from target to client
	if _, err := io.Copy(clientConn, targetConn); err != nil {
		//if err is connection reset by peer, we can ignore it
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			log.Printf("Connection reset by peer: %v", err)
			return
		}
		//if err is use of closed network connection, we can ignore it
		if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
			log.Printf("Temporary network error: %v", err)
			return
		}
		// Log other errors
		log.Printf("Error forwarding data from target to client: %v", err)
	}
}

func main() {

	var listenAddress string
	var connectAddress string

	flag.StringVar(&listenAddress, "listen", "", "Address to listen on (format: host:port or :port for all interfaces)")
	flag.StringVar(&connectAddress, "connect", "", "Target address to forward requests to (format: host:port)")
	flag.Parse()

	if listenAddress == "" {
		log.Fatal("Listen address must be specified in the format host:port or :port for all interfaces")
	}

	if connectAddress == "" {
		log.Fatal("Target address must be specified in the format host:port")
	}

	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", listenAddress, err)
	}
	defer listener.Close()

	log.Printf("Listening on %s, forwarding to %s", listenAddress, connectAddress)

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleConnection(clientConn, connectAddress)

		log.Printf("Accepted connection from %s", clientConn.RemoteAddr())
	}

}
