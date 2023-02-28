build:
	go build -o ./bin/mantenimiento

serve:
	./bin/mantenimiento

dk-build:
	docker build -t mantenimiento:alpine .

database:
	docker run --name postgres15 -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=Qwerty123 -d postgres:15.1-alpine

adminer:
	docker run --name adminer -p 8081:8080 -d adminer

rpc-desc:
	grpcurl -plaintext -msg-template localhost:5000 describe .GetBookListResponse

gen-proto:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.