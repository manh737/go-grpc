gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protos/*.proto
	
clean:
	rm protos/*.pb.go

test:
	go test -cover -race ./...

server:
	go run cmd/server/main.go

client:
	go run cmd/client/main.go