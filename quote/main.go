package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Here comes the embedding!
//go:embed quotes.txt
var quoteData string

func main() {
	rand.Seed(time.Now().Unix())
	quotes := strings.Split(quoteData, "\n\n")
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
