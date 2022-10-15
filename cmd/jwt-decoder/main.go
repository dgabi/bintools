/*
jwt-decoder decodes JWT strings and outputs to standard output

Usage:

	jwt-decoder eyJhbGciOiJIUzI1Ni......
*/
package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type decodedData map[string]interface{}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No JWT provided as input.")
		os.Exit(2)
	}

	parts := strings.Split(os.Args[1], ".")

	fmt.Println()
	fmt.Println(decode(parts[0]))
	fmt.Println(decode(parts[1]))
	fmt.Printf("Signature:%s\n", parts[2])
	fmt.Println()

}

func decode(s string) decodedData {
	d, err := b64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	var response decodedData

	if err := json.Unmarshal([]byte(d), &response); err != nil {
		panic(err)
	}
	return response
}
