package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/matheusjbatista/bergoglio/proto"
	"github.com/matheusjbatista/bergoglio/src/generateCode"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

type codeServer struct {
	pb.UnimplementedBergoglioServer
	mu sync.Mutex
}

func (s *codeServer) GetOneCodeWithBlessing(ctx context.Context, entry *pb.Entry) (*pb.End, error) {
	return &pb.End{Code: entry.GetName() + " Arvorinos garai"}, nil
}

type teste struct {
	arvorinho string
}

func (s *codeServer) GetManyCodes(ctx context.Context, manyCodesRequest *pb.ManyCodesRequest) (*pb.ManyCodesResponse, error) {
	fmt.Println("O cara entrou aqui")
	fmt.Println(manyCodesRequest)

	promotionalCodeSerialAndOrder := &generateCode.PromotionalCodeSerialAndOrder{
		SerialNumberInit:  manyCodesRequest.GetSerialNumberInit(),
		SerialNumberFinal: manyCodesRequest.GetSerialNumberFinal(),
		QuantityPerSerie:  manyCodesRequest.GetQuantityPerSerie(),
	}

	codes := generateCode.Generate(promotionalCodeSerialAndOrder)
	return &pb.ManyCodesResponse{Codes: codes}, nil
}

func (s *codeServer) GetManyCodesStream(manyCodesRequest *pb.ManyCodesRequest, stream pb.Bergoglio_GetManyCodesStreamServer) error {
	fmt.Println("Caiu aqui no stream em stream")
	promotionalCodeSerialAndOrder := &generateCode.PromotionalCodeSerialAndOrder{
		SerialNumberInit:  manyCodesRequest.GetSerialNumberInit(),
		SerialNumberFinal: manyCodesRequest.GetSerialNumberFinal(),
		QuantityPerSerie:  manyCodesRequest.GetQuantityPerSerie(),
	}

	codes := generateCode.Generate(promotionalCodeSerialAndOrder)
	start, toGet, toIncrease, isEndStream := 0, len(codes)/(len(codes)/100000), len(codes)/(len(codes)/100000), false

	for !isEndStream {
		if err := stream.Send(&pb.ManyCodesResponse{
			Codes: codes[start:toGet],
		}); err != nil {
			return err
		}

		start += toIncrease
		toGet += toIncrease

		if start == len(codes) {
			isEndStream = true
		}
	}

	return nil
}

func main() {
	port := 8563
	flag.Parse()
	listiner, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("Falhou ao subir o servidor: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterBergoglioServer(grpcServer, &codeServer{})
	fmt.Printf("Servidor rodando na porta: %v", port)
	grpcServer.Serve(listiner)
}
