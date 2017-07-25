// SHAHash
package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()
	count, error := h.Write([]byte("This needs to be hashed"))
	if error != nil {
		panic(error)
	}

	fmt.Println("No of bytes ", count)

	result := h.Sum(nil)

	fmt.Printf("%x\n", result)

}
