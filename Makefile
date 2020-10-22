all: anonreach.out

anonreach.out:
	protoc --go_out=. protobuf.proto
	go build -o bin/anonreach.out

clean:
	rm -rf bin/anonreach.out
	rm -rf utils/defs/protobuf.pb.go
