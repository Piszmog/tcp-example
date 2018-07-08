package main

import (
    "crypto/rsa"
    "crypto/rand"
    "log"
    "crypto/x509"
    "math/big"
    "crypto/x509/pkix"
    "time"
    "github.com/pkg/errors"
    "net"
    "encoding/pem"
    "fmt"
)

// Generation comes from http://pascal.bach.ch/2015/12/17/from-tcp-to-tls-in-go/ as a guide to creating private key
// and public certificate
func main() {
    rootKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        log.Panicln(errors.Wrap(err, "failed to generate key"))
    }

    rootCertTmpl, err := CertTemplate()
    if err != nil {
        log.Panicln(errors.Wrap(err, "failed to create cert template"))
    }
    rootCertTmpl.IsCA = true
    rootCertTmpl.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature
    rootCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}
    rootCertTmpl.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}
    _, rootCertPEM, err := CreateCert(rootCertTmpl, rootCertTmpl, &rootKey.PublicKey, rootKey)
    if err != nil {
        log.Panicln(errors.Wrap(err, "failed to create cert"))
    }
    fmt.Printf("%s\n", rootCertPEM)
    rootKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(rootKey),
    })
    fmt.Printf("%s\n", rootKeyPEM)
}

func CertTemplate() (*x509.Certificate, error) {
    // generate a random serial number (a real cert authority would have some logic behind this)
    serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
    serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
    if err != nil {
        return nil, errors.New("failed to generate serial number: " + err.Error())
    }

    tmpl := x509.Certificate{
        SerialNumber:          serialNumber,
        Subject:               pkix.Name{Organization: []string{"The world"}},
        SignatureAlgorithm:    x509.SHA256WithRSA,
        NotBefore:             time.Now(),
        NotAfter:              time.Now().Add(time.Hour), // valid for an hour
        BasicConstraintsValid: true,
    }
    return &tmpl, nil
}

func CreateCert(template, parent *x509.Certificate, pub interface{}, parentPriv interface{}) (
    cert *x509.Certificate, certPEM []byte, err error) {

    certDER, err := x509.CreateCertificate(rand.Reader, template, parent, pub, parentPriv)
    if err != nil {
        return
    }
    // parse the resulting certificate so we can use it again
    cert, err = x509.ParseCertificate(certDER)
    if err != nil {
        return
    }
    // PEM encode the certificate (this is a standard TLS encoding)
    b := pem.Block{Type: "CERTIFICATE", Bytes: certDER}
    certPEM = pem.EncodeToMemory(&b)
    return
}
