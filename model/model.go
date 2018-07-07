package model

type ComplexData struct {
    N int
    S string
    M map[string]int
    P []byte
    C *ComplexData
}