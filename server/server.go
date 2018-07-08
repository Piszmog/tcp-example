package server

import (
    "net"
    "log"
    "github.com/pkg/errors"
    "bufio"
    "crypto/tls"
)

type Handler interface {
    HandleMessages(readWriter *bufio.ReadWriter)
}

type Server struct {
    listener  net.Listener
    tlsConfig *tls.Config
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

func StartTLSServer(privateKey string, certificate string, handler Handler) {
    tlsConfig := createTLSConfig(certificate, privateKey)
    var tcpServer Server
    tcpServer.tlsConfig = tlsConfig
    err := tcpServer.Connect(":8081")
    defer tcpServer.Close()
    if err != nil {
        log.Panicln(errors.Wrap(err, "failed to connect the server"))
    }
    tcpServer.AcceptStreams(handler)
}

func createTLSConfig(certificate string, privateKey string) (*tls.Config) {
    cer, err := tls.X509KeyPair([]byte(certificate), []byte(privateKey))
    if err != nil {
        log.Panicln(errors.Wrap(err, "failed to create X509 key pair"))
    }
    tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}
    return tlsConfig
}

func (server *Server) Connect(address string) error {
    tlsConfig := server.tlsConfig
    var listener net.Listener
    var err error
    if tlsConfig != nil {
        listener, err = tls.Listen("tcp", address, tlsConfig)
    } else {
        listener, err = net.Listen("tcp", address)
    }
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
