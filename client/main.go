package main

import (
	"context"
	"dummy/httpconn"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	d := &dialer{httpconn.Dialer{InsecureSkipVerify: true}}

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithDialer(d.dial))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewSomeServiceClient(conn)
	r, err := client.MakeRPC(ctx, &SearchRequest{Query: "some query", PageNumber: 123, ResultPerPage: 456})
	if err != nil {
		panic(err)
	}

	fmt.Println("Got result: ", r.Result)
}

type dialer struct {
	httpconn.Dialer
}

func (d *dialer) dial(addr string, deadline time.Duration) (net.Conn, error) {
	return d.Dial(addr), nil
}
