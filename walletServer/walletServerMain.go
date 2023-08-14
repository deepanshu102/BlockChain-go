package main

import (
	"flag"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("wallet Server:")

}
func main() {
	port := flag.Uint("port", 8080, "TCP PORT number for wallet server")
	gateway := flag.String("gateway", "http://127.0.0.1:5000", "Blockchain Gateway")
	flag.Parse()
	fmt.Println(port)
	app := NewWalletServer(uint16(*port), *gateway)
	app.Run()

}
