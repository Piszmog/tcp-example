package main

import (
    "log"
    "bufio"
    "encoding/gob"
    "github.com/Piszmog/tcp-example/model"
    "github.com/Piszmog/tcp-example/server"
)

func main() {
    var handler GobHandler
    server.StartServer(&handler)
}

type GobHandler struct {
}

func (gobHandler *GobHandler) HandleMessages(readWriter *bufio.ReadWriter) {
    for {
        decoder := gob.NewDecoder(readWriter)
        var data model.ComplexData
        err := decoder.Decode(&data)
        if err != nil {
            log.Println("Reached EOF - close this connection")
            return
        }
        log.Printf("Outer complexData struct: \n%#v\n", data)
        log.Printf("Inner complexData struct: \n%#v\n", data.C)
    }
}
