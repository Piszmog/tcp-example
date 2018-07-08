package client

import (
    "net"
    "bufio"
    "crypto/x509"
    "log"
    "crypto/tls"
    "github.com/pkg/errors"
)

type Handler interface {
    SendMessages(readWriter *bufio.ReadWriter)
}

type Client struct {
    connection net.Conn
    tlsConfig  *tls.Config
}

func StartClient(handler Handler) {
    var client Client
    err := client.Connect("127.0.0.1:8081")
    defer client.Close()
    if err != nil {
        log.Panicln(errors.Wrap(err, "failed to connect thru tcp"))
    }
    client.writeMessages(handler)
}

func StartTLSClient(certificate string, handler Handler) {
    tlsConfig := createTLSConfig(certificate)
    var client Client
    client.tlsConfig = tlsConfig
    err := client.Connect("127.0.0.1:8081")
    defer client.Close()
    if err != nil {
        log.Panicln(errors.Wrap(err, "failed to connect thru tcp"))
    }
    client.writeMessages(handler)
}

func createTLSConfig(certificate string) *tls.Config {
    roots := x509.NewCertPool()
    ok := roots.AppendCertsFromPEM([]byte(certificate))
    if !ok {
        log.Panicln("failed to parse root certificate")
    }
    tlsConfig := &tls.Config{RootCAs: roots}
    return tlsConfig
}

func (client *Client) Connect(address string) error {
    var err error
    var conn net.Conn
    tlsConfig := client.tlsConfig
    if tlsConfig != nil {
        conn, err = tls.Dial("tcp", address, tlsConfig)
    } else {
        conn, err = net.Dial("tcp", address)
    }
    client.connection = conn
    return err
}

func (client *Client) Close() {
    defer client.connection.Close()
}

func (client *Client) writeMessages(handler Handler) {
    readWriter := bufio.NewReadWriter(bufio.NewReader(client.connection), bufio.NewWriter(client.connection))
    handler.SendMessages(readWriter)
    readWriter.Flush()
}
