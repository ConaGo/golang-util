package myutil

import (
	"bytes"
	cryptoRand "crypto/rand"
	"io"
	"math/rand"
	"time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func RandIntn(length int) int{
	return rand.Intn(length)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}


var rander = cryptoRand.Reader

// New generates a new uuid.
func NewUUID() ([16]byte, error) {
	var uuid [16]byte

	_, err := io.ReadFull(rander, uuid[:])
	if err != nil {
		return [16]byte{}, err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is 10

	return uuid, nil
}

// Equal returns true if two UUIDs are equal.
func Equal(a, b [16]byte) bool {
	return bytes.Equal([]byte(a[:]), []byte(b[:]))
}