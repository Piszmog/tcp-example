package server

import (
    "net"
    "log"
    "github.com/pkg/errors"
    "bufio"
)

type Handler interface {
    HandleMessages(readWriter *bufio.ReadWriter)
}

type Server struct {
    listener net.Listener
}

func StartServer(handler Handler) {
    var tcpServer Server
    err := tcpServer.Connect(":8081")
    defer tcpServer.Close()
    if err != nil {
        log.Panicln(errors.Wrap(err, "failed to connect the server"))
    }
    tcpServer.AcceptStreams(handler)
}

func (server *Server) Connect(address string) error {
    listener, err := net.Listen("tcp", address)
    server.listener = listener
    return err
}

func (server *Server) Close() {
    server.listener.Close()
}

func (server *Server) AcceptStreams(handler Handler) {
    for {
        conn, err := server.listener.Accept()
        if err != nil {
            log.Panicln(errors.Wrap(err, "failed to accept TCP"))
        }
        go handleConnection(conn, handler)
    }
}

func handleConnection(conn net.Conn, handler Handler) {
    defer conn.Close()
    readWriter := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
    handler.HandleMessages(readWriter)
}
