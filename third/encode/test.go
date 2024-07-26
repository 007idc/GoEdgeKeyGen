package encode

import (
	"fmt"
	"log"
)

func Decode(data []byte) []byte {
	method, err := NewMethodInstance("aes-256-cfb", "41100c93a65cfb71d5b0672c0d60d7ec", "70ba69d67bf7e61e17ac565c6093a325")
	if err != nil {
		fmt.Println("[MagicKeyEncode]" + err.Error())
		return data
	}

	src, err := method.Decrypt(data)
	if err != nil {
		fmt.Println("[MagicKeyEncode]" + err.Error())
		return data
	}
	return src
}
func Encode(data []byte) []byte {
	method, err := NewMethodInstance("aes-256-cfb", "41100c93a65cfb71d5b0672c0d60d7ec", "70ba69d67bf7e61e17ac565c6093a325")
	if err != nil {
		fmt.Println("[MagicKeyEncode]" + err.Error())
		return data
	}

	dst, err := method.Encrypt(data)
	if err != nil {
		fmt.Println("[MagicKeyEncode]" + err.Error())
		return data
	}
	return dst
}

func Test() {
	method, err := NewMethodInstance("aes-256-cfb", "abc", "123")
	if err != nil {
		log.Fatal(err)
	}
	src := []byte("Hello, World")
	dst, err := method.Encrypt(src)
	if err != nil {
		log.Fatal(err)
	}
	dst = dst[:len(src)]
	fmt.Println("dst:", string(dst))

	src, err = method.Decrypt(dst)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("src:", string(src))
}
