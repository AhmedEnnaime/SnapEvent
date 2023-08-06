gen:
	protoc --proto_path=proto proto/*.proto  --go_out=:pb --go-grpc_out=:pb
db:
	docker-compose -f ./docker/docker-compose.yml up -d
db-down:
	docker-compose -f ./docker/docker-compose.yml down