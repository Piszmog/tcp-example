# TCP Server Example
An example of a TCP server and TCP client written in Go.

I wrote two examples of using TCP. One server/client is using just string to send messages. These string end with the newline
character `\n`. The other server/client is streaming [Gob](https://golang.org/pkg/encoding/gob/) data using `ComplexData` model.

## Server
The server API resides in the `server` package.

To start a server, create a handler that implements the interface `server.Handler` and run `server.StartServer(..)`

Example,
```go
func main() {
    var handler ExampleHandler
    server.StartServer(&handler)
}

type StringHandler struct {
}

func (handler *ExampleHandler) HandleMessages(readWriter *bufio.ReadWriter) {
    ...
}
```

## Client
The client API resides in the `client` package.

To create a client and to send data to the server, a handler that implements `client.Handler` and run `client.StartClient(..)`

Example,
```go
func main() {
    var handler ExampleClientHandler
    client.StartClient(&handler)
}

type ExampleClientHandler struct {
}

func (handler *ExampleClientHandler) SendMessages(readWriter *bufio.ReadWriter) {
    ...
}
```

## TLS Server/Client
To create a secure server and to allow a client to connect to it, the `certgenerator` can be ran to generate a private key and 
a public certificate.

To use the private key and public certificate, instead of using `server.StartServer(..)` and `client.StartClient(..)` use 
`server.StartTLSServer(..)` and `client.StartTLSClient(..)`.

## References
* https://appliedgo.net/networking/
  * insecure tcp server and client
* http://pascal.bach.ch/2015/12/17/from-tcp-to-tls-in-go/
  * tls server and client
* https://ericchiang.github.io/post/go-tls/
  * generation of private key and public cert