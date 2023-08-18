gen:
	mkdir -p pb
	protoc --proto_path=proto proto/*.proto --go_out=:pb --go-grpc_out=:pb

up:
	docker-compose up -d

build:
	docker-compose up -d --build
	
down:
	docker-compose down

clean:
	rm pb/*.go

test:
	go test -cover -race ./...

server:
	go run cmd/server/main.go -port 8080

clients:
	go run cmd/client/main.go -address 0.0.0.0:8080
