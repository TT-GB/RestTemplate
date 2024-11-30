# TT Rest Example

Basic repository to allow the creation of containerized rest api endpoints

Basic automated testing included


# Steps
go mod init restapi
go mod tidy

Run DockerDesktop
docker build -t go-rest-api .


docker run -p 8080:8080 go-rest-api
or 
docker-compose up


then visit or use curl on http://localhost:8080/api/hello

# Running tests

locally
go test -v ./...

Through docker (it will run the tests as part of the build in this example)
docker build -t go-rest-api .

# Prereqs
go install google.golang.org/grpc@latest

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/
