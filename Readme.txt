Step 1: make a proto file in a folder named proto.

Step 2: install protoc-gen-go and protoc-gen-go-grpc packages using:

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest	

Step 3: make a mod file using: 
	
	go mod init anyNameYouLike

Step 4: add both installed packages to mod file using: 

	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

Step 5: generate Go code using:

	protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=proto --go-grpc_opt=paths=source_relative protoFileName.proto 

	this command will generate two files, one will have code for all the messages(protoFileName.pb.go) and the second will have all the code for the services(protoFileName_grpc.pb.go).

Step 6: to get all the missing packages use:
	
	go mod tidy
	
	this will automatically get all the missing packages. 	
