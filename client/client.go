package main

import (
	"context"
	"fmt"
	pb "grpc-playground/protos/people"

	"google.golang.org/grpc"
)

func main() {
	serverAddress := "127.0.0.1:3000"

	conn, _ := grpc.Dial(serverAddress, grpc.WithInsecure())

	peopleClient := pb.NewPeopleInfoServiceClient(conn)

	//var wg sync.WaitGroup
	//wg.Add(1)

	var reqs = []*pb.PeopleInfoRequest{
		{
			Name:    "Jeremy",
			Address: "Taipei",
			Age:     33,
			Hobby:   "sleep",
			Height:  180,
			Weight:  80,
		},
		{
			Name:    "Vera",
			Address: "Taipei",
			Age:     34,
			Hobby:   "sleep",
			Height:  156,
			Weight:  52,
		},
	}

	stream, err := peopleClient.PeopleInfo(context.Background())
	if err != nil {
		fmt.Println("Cannot create stream: ", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		res, _ := stream.CloseAndRecv()
		fmt.Println(res)
	}

	//wg.Wait()
}
