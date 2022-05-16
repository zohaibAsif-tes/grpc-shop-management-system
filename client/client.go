package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zohaibAsif-tes/grpc-shop-management-system/proto"
	"google.golang.org/grpc"
)

func main() {
	//creating a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //opening an insecure connection
	if err != nil {
		log.Fatalf("Client Error :: Cannot Connect : %v\n", err)
	}
	fmt.Println("Connected to the server successfully.")
	//defering closing the connection untill finished
	defer conn.Close()

	client := proto.NewShopManagementServiceClient(conn)

	//creating a gin server
	g := gin.Default()

	//sending request to server and getting response from server
	g.POST("/generateBill", func(ctx *gin.Context) {

		ctx.Header("Content-Type", "application/json")

		bill := &proto.BillRequest{}

		err := json.NewDecoder(ctx.Request.Body).Decode(&bill)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			log.Fatalf("Client Error :: Cannot Decode Request Body : %v\n", err)
		}

		billRequest := &proto.BillRequest{
			Bill: bill.GetBill(),
		}

		response, err := client.GenerateBill(ctx, billRequest)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			log.Fatalf("Client Error :: Cannot Get Response from Server : %v\n", err)
		}

		json.NewEncoder(ctx.Writer).Encode(response)
	})

	//runnning gin server
	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Client Error :: Failed to Run Gin Server: %v", err)
	}
}
