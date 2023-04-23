package main

import (
	"fmt"
	"os"
	"sort"
	"strings")

func main(){
	path := os.Getenv("PATH")
	if path == "" {
		fmt.Println("PATH not found")
		os.Exit(-1)
	}
	values := strings.SplitN(path,":",-1)
	sort.Strings(values)
	for _,v := range values {
		fmt.Println(v)
	}
}
