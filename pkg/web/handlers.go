package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	bchain "github.com/chespinoza/toy-blockchain-go/bchain"
	"github.com/davecgh/go-spew/spew"
)

type Message struct {
	BPM int
}

// BlockChain ...
var BlockChain []bchain.Block

func writeBlockChainHandler(w http.ResponseWriter, r *http.Request) {
	var m Message
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()
	newBlock, err := bchain.GenerateBlock(BlockChain[len(BlockChain)-1], m.BPM)
	if err != nil {
		respondWithJSON(w, r, http.StatusInternalServerError, m)
		return
	}
	if bchain.ValidateBlock(newBlock, BlockChain[len(BlockChain)-1]) {
		newBlockChain := append(BlockChain, newBlock)
		bchain.ReplaceChain(newBlockChain, &BlockChain)
		spew.Dump(BlockChain)
	} else {
		log.Println("attempt to add non valid block!")
		respondWithJSON(w, r, http.StatusInternalServerError, nil)
		return
	}
	respondWithJSON(w, r, http.StatusCreated, newBlock)
}

func getBlockChainHandler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(BlockChain, "", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
