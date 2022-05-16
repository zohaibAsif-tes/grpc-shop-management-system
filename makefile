gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/shop_service.proto

remove:
	rm proto/*.go	

run server:
	go run server/server.go

run client:
	go run client/client.go
