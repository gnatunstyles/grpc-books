.PHONY: gen
gen:
	protoc --go_out=./ --go_opt=paths=import --go-grpc_out=./ --go-grpc_opt=paths=import proto/app.proto 

.PHONY: build
build:
	go build -v ./cmd/server/

.PHONY: run
run:
	./server