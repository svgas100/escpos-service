

run: build
	# run
	export GRPC_GO_LOG_VERBOSITY_LEVEL=99
	export GRPC_GO_LOG_SEVERITY_LEVEL=info

	go run -v ./cmd/escpos-service.go

build: clean
	# generate proto files
	protoc --go_out=./internal/grpc --go-grpc_out=./internal/grpc ./api/escpos-service.proto 

	# compile dependencies
	go build -v ./...
	
	# create executable
	go build -v -o escpos-service.so ./cmd/escpos-service.go

clean:
	# clean
	rm -f -R ./internal/grpc/*
	rm -f escpos-service.so