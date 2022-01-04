package main

import (
	"github.com/639dede/ydcoin/cli"
	"github.com/639dede/ydcoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
