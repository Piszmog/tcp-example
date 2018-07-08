package main

import (
    "github.com/Piszmog/tcp-example/client"
    "bufio"
)

func main() {
    var handler StringClientHandler
    // connecting to an insecure server
    client.StartClient(&handler)
    // connecting to a secure server
    //client.StartTLSClient(cert.PublicCertificate, &handler)
}

type StringClientHandler struct {
}

func (handler *StringClientHandler) SendMessages(readWriter *bufio.ReadWriter) {
    readWriter.WriteString("Hello\n")
    readWriter.WriteString("Again\n")
}