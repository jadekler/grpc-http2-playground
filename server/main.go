package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	RegisterSomeServiceServer(s, &server{})

	fmt.Println("Serving on 8080")
	err = s.Serve(l)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}

type server struct{}

func (s *server) MakeRPC(ctx context.Context, r *SearchRequest) (*SearchResult, error) {
	fmt.Println("Got query: ", r.Query)

	return &SearchResult{Query: r.Query, PageNumber: r.PageNumber, Result: "some result"}, nil
}
