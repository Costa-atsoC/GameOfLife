BINARY_NAME=main

run: 
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME} -s=100 -t=150 -x=20 -y=15

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

clean:
	go clean