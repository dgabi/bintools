package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Symbol int

const (
	Letters  Symbol = 0
	Numbers  Symbol = 1
	Specials Symbol = 2
)

var odds = [10]Symbol{1, 1, 2, 2, 0, 0, 0, 0, 0, 0}

func main() {
	rand.Seed(time.Now().UnixNano())
	maxlen := 16
	symbols := map[Symbol]string{
		Letters:  "qwertyuiopasdfghjklzxcvbnm",
		Numbers:  "0123456789",
		Specials: "~!@#$%^&*()_+",
	}

	str := ""
	for i := 1; i <= maxlen; i++ {
		kind := GetNextType()
		c := PickSymbol(symbols[kind])
		str = str + c
	}
	fmt.Println(str)
}

func GetNextType() Symbol {
	var r int = rand.Intn(10)
	return Symbol(odds[r])
}

func PickSymbol(s string) string {
	pos := rand.Intn(len(s))
	return s[pos : pos+1]
}
