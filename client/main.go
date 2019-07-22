package main

import (
    "flag"
    "google.golang.org/grpc"
    "log"
    pt "../proto"
    "golang.org/x/net/context"
    "time"
)

var client pt.BlockchainClient

func main() {
    addFlag := flag.Bool("add", false, "add new Block")
    listFlag := flag.Bool("list", false, "get the blockchain")
    flag.Parse()

    conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("cannot dial server: %v", err)
    }

    client = pt.NewBlockchainClient(conn)

    if *addFlag {
        addBlock()
    }
    if *listFlag {
        getBlockchain()
    }
}

func addBlock() {
    block, err := client.AddBlock(context.Background(), &pt.AddBlockRequest{
        Data: time.Now().String(),
    })
    if err != nil {
        log.Fatalf("unable to add block: %v", err)
    }
    log.Printf("new block hash: %s\n", block.Hash)
}

func getBlockchain() {
    bc, err := client.GetBlockChain(context.Background(), &pt.GetBlockchainRequest{})
    if err != nil {
        log.Fatalf("unable to get blockchain: %v", err)
    }
    log.Printf("blocks:")
    for _, b := range bc.Blocks {
        log.Printf("hash: %s, prev block hash: %s, data: %s", b.Hash, b.PrevBlockHash, b.Data)
    }
}
