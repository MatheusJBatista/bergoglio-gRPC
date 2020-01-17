package main

import (
	"context"
	"fmt"
	pb "github.com/matheusjbatista/bergoglio/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func printOneCodeWithBlessing(client pb.BergoglioClient, entry *pb.Entry) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	codeWithBlessing, err := client.GetOneCodeWithBlessing(ctx, entry)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(codeWithBlessing)
}

func printManyCodes(client pb.BergoglioClient, manyCodesRequest *pb.ManyCodesRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	codes, err := client.GetManyCodes(ctx, manyCodesRequest)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(codes)
}

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	serverConnection, err := grpc.Dial("localhost:58147", opts...)

	if err != nil {
		log.Fatalf("Erro ao comunicar com servidor: %v", err)
	}
	defer serverConnection.Close()

	client := pb.NewBergoglioClient(serverConnection)

	//printOneCodeWithBlessing(client, &pb.Entry{Name: "Matheus Batista"})
	printManyCodes(client, &pb.ManyCodesRequest{
		SerialNumberInit:  88,
		SerialNumberFinal: 99,
		QuantityPerSerie:  44,
	})
}
