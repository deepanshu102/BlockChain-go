package main

import (
	"flag"
	"log"
)

func init() {
	log.SetPrefix("wallet Server:")

}
func main() {
	port := flag.Uint("port", 8080, "TCP PORT number for wallet server")
	gateway := flag.String("gateway", "http://127.0.0.1:5000", "Blockchain Gateway")
	flag.Parse()
	app := NewWalletServer(uint16(*port), *gateway)
	app.Run()

}
