all: anonreach.out

anonreach.out:
	go build -o bin/anonreach.out

clean:
	rm -rf bin/anonreach.out
