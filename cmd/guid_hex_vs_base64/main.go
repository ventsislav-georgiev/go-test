package main

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
	"github.com/jxskiss/base62"
)

func guidHex() string {
	return uuid.New().String()
}

func guidBase64() string {
	b, _ := uuid.New().MarshalBinary()
	s := base64.URLEncoding.EncodeToString(b)
	return s[:len(s)-2]
}

func guidBase62() string {
	b, _ := uuid.New().MarshalBinary()
	return base62.EncodeToString(b)
}

func main() {
	fmt.Println(guidBase62())
	fmt.Println(len(guidBase62()))
}
