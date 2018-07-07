package main

import (
    "bufio"
    "github.com/Piszmog/tcp-example/client"
)

type StringClientHandler struct {
}

func (handler *StringClientHandler) SendMessages(readWriter *bufio.ReadWriter) {
    readWriter.WriteString("Hello\n")
    readWriter.WriteString("Again\n")
}

func main() {
    var handler StringClientHandler
    client.StartClient(&handler)
}
