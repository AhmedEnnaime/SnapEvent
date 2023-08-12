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
