package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	pb "github.com/knative-sample/grpc-stream/proto"
	"google.golang.org/grpc"
)

func main() {

	countStr := os.Getenv("GRPC_CONCURRENT")
	serverAddr := os.Getenv("GRPC_SERVER_ADDR")
	conn, err := grpc.Dial(
		serverAddr,
		[]grpc.DialOption{
			grpc.WithInsecure(),
			//grpc.WithTimeout(time.Second * 300),
		}...,
	)
	if err != nil {
		log.Fatalf("grpc.Dial error: %s", err)
	}

	defer conn.Close()

	client := pb.NewStreamServiceClient(conn)

	count, _ := strconv.Atoi(countStr)
	if count == 0 {
		count = 1000
	}
	wg := sync.WaitGroup{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(index int) {
			err = printStreamLists(
				client,
				&pb.StreamRequest{
					Msg: &pb.StreamMessage{
						Key:   fmt.Sprintf("Client Stream List for [%d]", index),
						Value: int32(index),
					},
				},
			)
			if err != nil {
				log.Printf("printStreamLists error: %s", err)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func printStreamLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.StreamList(context.Background(), r)
	if err != nil {
		return err
	}

	stime := time.Now()
	var resp *pb.StreamResponse
	for {
		_resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		resp = _resp
	}
	log.Printf("resp: key: %s, latency:%d", resp.Msg.Key, time.Now().Unix()-stime.Unix())

	stream.CloseSend()
	return nil
}
