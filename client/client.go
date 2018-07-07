package client

import (
    "net"
    "bufio"
)

type Handler interface {
    SendMessages(readWriter *bufio.ReadWriter)
}

type Client struct {
    connection net.Conn
}

func StartClient(handler Handler) {
    var client Client
    client.Connect("127.0.0.1:8081")
    defer client.Close()
    client.writeMessages(handler)
}

func (client *Client) Connect(address string) error {
    conn, err := net.Dial("tcp", address)
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
