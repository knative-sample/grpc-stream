package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	pb "github.com/knative-sample/grpc-stream/proto"
	"google.golang.org/grpc"
)

type StreamService struct{}

const (
	Port = "8080"
)

func main() {
	maxConcurrentStreams, _ := strconv.Atoi(os.Getenv("GRPC_MAX_CONCURRENT_STREAMS"))
	server := grpc.NewServer(grpc.MaxConcurrentStreams(uint32(maxConcurrentStreams)))
	pb.RegisterStreamServiceServer(server, &StreamService{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("net.Listen error: %s", err)
	}

	server.Serve(lis)
}

func (s *StreamService) StreamList(r *pb.StreamRequest, stream pb.StreamService_StreamListServer) error {
	for i := 1; i < 2; i++ {
		err := stream.Send(&pb.StreamResponse{
			Msg: &pb.StreamMessage{
				Key:   fmt.Sprintf("%s  --> resp index:%d", r.Msg.Key, r.Msg.Value),
				Value: int32(i),
			},
		})
		if err != nil {
			log.Printf("StreamList error:%s", err.Error())
			return err
		}
		time.Sleep(time.Second * 2)
	}

	return nil
}
