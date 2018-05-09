package main

import (
	"crypto/sha256"
	"fmt"
	"encoding/binary"
	"math"
)


func uint_tobytearrs(x uint32) []byte{
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint32(bs, x)//change int to byte array
	//fmt.Println(bs)
	return bs
}

func Float64frombyte(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func main() {
	bs:=uint_tobytearrs(uint32(100000000))
	x1:=sha256.Sum256(bs)
	x:=x1[:]
	fmt.Println(math.Abs(Float64frombyte(x)))
}
