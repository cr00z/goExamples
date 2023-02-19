.PHONY: client cert

gen: clean
	#protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative proto/*.proto
	protoc --proto_path=proto --go_out=pb \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto

clean:
	-@rm pb/*.go

server:
	go run cmd/server/main.go -port 8080

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

nginx:
	cd deploy/nginx; docker compose up -d

ndown:
	cd deploy/nginx; docker compose down

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...

cert:
	cd cert; chmod +x gen.sh; ./gen.sh; cd ..