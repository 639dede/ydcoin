package cli

import (
	"flag"
	"fmt"
	"os"

	explorer "github.com/639dede/ydcoin/explorer/templates"
	"github.com/639dede/ydcoin/rest"
)

func usage(){
	fmt.Printf("Welcome to ydcoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port1 : Set the PORT of the rest API server\n")
	fmt.Printf("-port2 : Set the PORT of the html server\n")
	os.Exit(0)
}

func Start(){
	if len(os.Args) < 2{
		usage()
	}
	port1 := flag.Int("port1", 3000, "Set port of the rest API server")

	port2 := flag.Int("port2", 5000, "Set port of the html server")

	flag.Parse()

	go rest.Start(*port1)
	explorer.Start(*port2)
}