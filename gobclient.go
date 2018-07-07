package main

import (
    "bufio"
    "encoding/gob"
    "github.com/Piszmog/tcp-example/model"
    "github.com/Piszmog/tcp-example/client"
)

type GobClientHandler struct {
}

func (handler *GobClientHandler) SendMessages(readWriter *bufio.ReadWriter) {
    testStruct := model.ComplexData{
        N: 23,
        S: "string data",
        M: map[string]int{"one": 1, "two": 2, "three": 3},
        P: []byte("abc"),
        C: &model.ComplexData{
            N: 256,
            S: "Recursive structs? Piece of cake!",
            M: map[string]int{"01": 1, "10": 2, "11": 3},
        },
    }
    encoder := gob.NewEncoder(readWriter)
    encoder.Encode(testStruct)
}

func main() {
    var handler GobClientHandler
    client.StartClient(&handler)
}
