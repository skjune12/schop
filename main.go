package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/skjune12/schop/lib"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("Argument length must be 1.")
	}

	r := schop.Search(args[0])
	bytes, _ := r.ToJson()
	fmt.Println(bytes)
}
