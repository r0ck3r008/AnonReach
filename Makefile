all: anonreach.out

anonreach.out:
	protoc --go_out=. utils/protobuf.proto
	go build -o bin/anonreach.out

clean:
	rm -rf bin/anonreach.out
	rm -rf utils/protobuf.pb.go
