package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/zohaibAsif-tes/grpc-shop-management-system/database"
	"github.com/zohaibAsif-tes/grpc-shop-management-system/models"
	"github.com/zohaibAsif-tes/grpc-shop-management-system/proto"
	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedShopManagementServiceServer
}

func main() {
	//creating a listener to listen on the default gRPC port(50051)
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Server Error :: Failed to Listen : %v\n", err)
	}

	//creating a new gRPC server
	server := grpc.NewServer()
	proto.RegisterShopManagementServiceServer(server, &Server{})

	fmt.Println("Starting Server...")

	//binding the port to gRPC server
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Server Error :: Failed to Serve : %v\n", err)
	}
}

func init() {
	//establishing a connection with the database
	database.EstablishConnection()
}

func (s *Server) GenerateBill(ctx context.Context, bill *proto.BillRequest) (*proto.BillResponse, error) {

	fmt.Printf("GenerateBill function was invoked with : %v\n", bill)

	//getting list of products from request
	listOfProducts := bill.GetBill().GetListOfProducts()

	//calculating total bill amount
	var total float32
	for i := 0; i < len(listOfProducts); i++ {
		price := listOfProducts[i].GetPrice()
		total += price
	}

	//concatenating customer's first and last name
	customerName := bill.GetBill().GetCustomer().GetFirstName() + bill.GetBill().GetCustomer().GetLastName()
	//creating an object of Bill model to store in the database
	generatedBill := &models.Bill{Name: customerName, Total: total}

	//saving bill in database
	database.DB.Save(generatedBill)

	// generating a response of type proto.BillResponse{}
	response := &proto.BillResponse{
		Id:        int32(generatedBill.ID),
		Total:     total,
		Bill:      bill.GetBill(),
		CreatedAt: generatedBill.CreatedAt.String(),
	}

	//returning response to grpc client
	return response, nil
}
