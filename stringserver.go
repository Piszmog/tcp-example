package main

import (
    "log"
    "bufio"
    "io"
    "strings"
    "github.com/Piszmog/tcp-example/server"
)

func main() {
    var handler StringHandler
    // creating an insecure server
    server.StartServer(&handler)
    // creating a secure server
    //server.StartTLSServer(cert.PrivateKey, cert.PublicCertificate, &handler)
}

type StringHandler struct {
}

func (stringHandler *StringHandler) HandleMessages(readWriter *bufio.ReadWriter) {
    for {
        message, err := readWriter.ReadString('\n')
        switch {
        case err == io.EOF:
            log.Println("Reached EOF - close this connection")
            return
        case err != nil:
            log.Println("\nError reading command. Got: "+message, err)
            return
        }
        message = strings.Trim(message, "\n ")
        log.Println("Recieved " + message)
    }
}
