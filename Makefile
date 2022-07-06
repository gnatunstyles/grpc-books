# SHELL := /bin/bash


.PHONY: gen
gen:
	protoc --go_out=api --go_opt=paths=import --go-grpc_out=api --go-grpc_opt=paths=import proto/app.proto 