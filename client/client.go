package main

import (
	"context"
	"fmt"
	pb "github.com/matheusjbatista/bergoglio/proto"
	"google.golang.org/grpc"
	"io"
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

func printManyCodesStream(client pb.BergoglioClient, manyCodesRequest *pb.ManyCodesRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	codesStream, err := client.GetManyCodesStream(ctx, manyCodesRequest)

	if err != nil {
		log.Fatal(err)
	}

	for {
		streamRecovery, err := codesStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Stream data: ", streamRecovery.Codes)
	}
}

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	serverConnection, err := grpc.Dial("localhost:8563", opts...)

	if err != nil {
		log.Fatalf("Erro ao comunicar com servidor: %v", err)
	}
	defer serverConnection.Close()

	client := pb.NewBergoglioClient(serverConnection)

	manyCodesRequest := &pb.ManyCodesRequest{
		SerialNumberInit:  88,
		SerialNumberFinal: 99,
		QuantityPerSerie:  90000,
	}

	printManyCodesStream(client, manyCodesRequest)

	//printOneCodeWithBlessing(client, &pb.Entry{Name: "Matheus Batista"})
	// printManyCodes(client, &pb.ManyCodesRequest{
	// 	SerialNumberInit:  88,
	// 	SerialNumberFinal: 99,
	// 	QuantityPerSerie:  900,
	// })
}
