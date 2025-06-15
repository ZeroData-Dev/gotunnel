package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"time"
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
	logfilename := "gotunnel_" + time.Now().Format("20060102_150405") + ".log"
	logFile, err := os.OpenFile(filepath.Join(os.TempDir(), logfilename), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	log.Printf("Logging to %s", logFile.Name())
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))

	var listenAddress string
	var connectAddress string

	flag.StringVar(&listenAddress, "listen", "", "Address to listen on (format: host:port or :port for all interfaces)")
	flag.StringVar(&connectAddress, "connect", "", "Target address to forward requests to (format: host:port or :port for all interfaces)")
	flag.Parse()

	if listenAddress == "" || connectAddress == "" {
		flag.Usage()
		return
	}

	listener, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("Failed to listen on %s: \n%v\n", listenAddress, err)
	}
	defer listener.Close()

	log.Printf("Listening on %s, forwarding to %s\n", listenAddress, connectAddress)

	for {
		clientConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: \n%v\n", err)
			continue
		}

		go handleConnection(clientConn, connectAddress)

		log.Printf("Accepted connection from %s\n", clientConn.RemoteAddr())
	}

}
