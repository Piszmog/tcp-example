package main

import (
    "bufio"
    "github.com/Piszmog/tcp-example/client"
)

type StringClient struct {
}

func (stringClient *StringClient) SendMessages(readWriter *bufio.ReadWriter) {
    readWriter.WriteString("Hello\n")
    readWriter.WriteString("Again\n")
}

func main() {
    var handler StringClient
    client.StartClient(&handler)
}
