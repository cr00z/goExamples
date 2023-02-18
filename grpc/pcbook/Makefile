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

servertls:
	go run cmd/server/main.go -port 8080 -tls=true

client:
	go run cmd/client/main.go

test:
	go test -cover -race ./...

cert:
	cd cert; chmod +x gen.sh; ./gen.sh; cd ..