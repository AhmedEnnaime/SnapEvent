gen:
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb
up:
	docker-compose up -d
down:
	docker-compose down