package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/diltheyaislan/grpc-golang/pb/pb"
	"google.golang.org/grpc"
)

func main() {

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	fmt.Println("Select function:\n",
		"  1 - AddUser\n",
		"  2 - AddUserVerbose\n",
		"  3 - AddUsers")

	reader := bufio.NewReader(os.Stdin)
	option, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	switch option {
	case '1':
		fmt.Println("AddUser was selected")
		AddUser(client)
		break
	case '2':
		fmt.Println("AddUserVerbose was selected")
		AddUserVerbose(client)
		break
	case '3':
		fmt.Println("AddUsers was selected")
		AddUsers(client)
		break
	}
}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Dilthey Aislan",
		Email: "dilthey@aislan.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Dilthey Aislan",
		Email: "dilthey@aislan.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)
		}

		fmt.Println("Status:", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "u1",
			Name:  "Dilthey Aislan",
			Email: "dilthey@aislan.com",
		},
		&pb.User{
			Id:    "u2",
			Name:  "Noah de Paula",
			Email: "noah@noah.com",
		},
		&pb.User{
			Id:    "u3",
			Name:  "Aislan Dilthey",
			Email: "aislan@dilthey.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Could not create request", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not receive response", err)
	}

	fmt.Println(res)
}
