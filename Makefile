run: build
	# run 
	go run -v ./cmd/escpos-service.go

build:
	# compile dependencies
	go build -v ./...
	
	# create executable
	go build -v -o escpos-service.so ./cmd/escpos-service.go

