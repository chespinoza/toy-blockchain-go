package main

import (
	"log"
	"os"
	"time"

	"github.com/chespinoza/toy-blockchain-go/pkg/web"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	httpAddr := os.Getenv("ADDR")

	go func() {
		t := time.Now()
		initialBlock := Block{0, t.String(), 0, "", ""}
		spew.Dump(initialBlock)
		BlockChain = append(BlockChain, initialBlock)
	}()
	log.Fatal(web.Run(httpAddr))
}
