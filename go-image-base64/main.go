package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// open file
	f, err := os.Open("ktp.jpg")
	if err != nil {
		panic(err)
	}

	// read jpg into byte
	reader := bufio.NewReader(f)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	// encode as base64
	encoded := base64.StdEncoding.EncodeToString((content))

	fmt.Println("ENCODED: ", encoded)

}
