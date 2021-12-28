package main

import (
	explorer "github.com/639dede/ydcoin/explorer/templates"
	"github.com/639dede/ydcoin/rest"
)

func main(){
	go rest.Start(4000)
	explorer.Start(3000)
}