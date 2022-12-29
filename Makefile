.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	go install github.com/bradleyjkemp/grpc-tools/...@latest
	go get github.com/Syfaro/telegram-bot-api@latest

.PHONY: proto
proto:
	rm -fr pb
	mkdir -p pb

	protoc -I api/proto --go_out=pb --go_opt=paths=source_relative \
           --go-grpc_out=pb --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
           --grpc-gateway_out=pb  --grpc-gateway_opt=logtostderr=true,paths=source_relative,generate_unbound_methods=true \
            tg.proto
