package main

import (
	"crypto/rand"
	"fmt"
)

func FakeGenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
