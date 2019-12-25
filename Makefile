all: pb/hello.pb.go server/server client/client
.PHONY: all

pb/hello.pb.go: pb/hello.proto
	protoc -I pb pb/hello.proto --go_out=plugins=grpc:pb

server/server: server/server.go
	cd server && go build server.go

client/client: client/client.go
	cd client && go build client.go

clean:
	rm pb/hello.pb.go
	rm server/server
	rm client/client

run:
	./server/server &
	./client/client
