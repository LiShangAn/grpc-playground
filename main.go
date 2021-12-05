package main

import (
	"fmt"
	pb "grpc-playground/protos/people"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type service struct {
	pb.UnimplementedPeopleInfoServiceServer
}

func (s *service) PeopleInfo(stream pb.PeopleInfoService_PeopleInfoServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("no more data")
			break
		}
		if err != nil {
			return fmt.Errorf("error when receive stream request: %v", err)
		}
		name := req.GetName()
		hobby := req.GetHobby()
		addess := req.GetAddress()
		age := req.GetAge()
		height := req.GetHeight()
		weight := req.GetWeight()

		fmt.Printf("%s, %s, %s, %d, %d, %d\n", name, hobby, addess, age, height, weight)
	}

	res := &pb.PeopleInfoResponse{
		Status:  true,
		Message: "success",
	}
	if err := stream.SendAndClose(res); err != nil {
		return fmt.Errorf("sending response fail: %v", err)
	}

	return nil
}

func main() {

	addr := "0.0.0.0:3000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Server listening on", addr)
	gRPCServer := grpc.NewServer()

	pb.RegisterPeopleInfoServiceServer(gRPCServer, &service{})
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
