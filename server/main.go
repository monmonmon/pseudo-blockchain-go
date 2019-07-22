package main

import (
    pt "../proto"
    "./blockchain"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "log"
    "net"
)

type Server struct {
    Blockchain blockchain.Blockchain
}

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("unable to listen on 8080 port: %v", err)
    }

    srv := grpc.NewServer()
    bc := blockchain.NewBlockchain()
    pt.RegisterBlockchainServer(srv, &Server{
        Blockchain: *bc,
    })
    srv.Serve(listener)
}

func (s *Server) AddBlock(ctx context.Context, in *pt.AddBlockRequest) (*pt.AddBlockResponse, error) {
    block := s.Blockchain.AddBlock(in.Data)
    return &pt.AddBlockResponse{
        Hash: block.Hash,
    }, nil
    // return new(pt.AddBlockResponse), nil
}

func (s *Server) GetBlockChain(ctx context.Context, in *pt.GetBlockchainRequest) (*pt.GetBlockchainResponse, error) {
    resp := new(pt.GetBlockchainResponse)
    for _, b := range s.Blockchain.Blocks {
        resp.Blocks = append(resp.Blocks, &pt.Block{
            PrevBlockHash: b.PrevBlockHash,
            Hash: b.Hash,
            Data: b.Data,
        })
    }
    return resp, nil
    // return new(pt.GetBlockchainResponse), nil
}
