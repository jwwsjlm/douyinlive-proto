GO ?= go
MODULE := github.com/jwwsjlm/douyinlive-proto

.PHONY: proto test tidy clean help

proto:
	@echo "Generating Go code from protobuf/new_douyin.proto..."
	protoc --proto_path=protobuf --go_out=. --go_opt=module=$(MODULE) protobuf/new_douyin.proto

test:
	$(GO) test ./...

tidy:
	$(GO) mod tidy

clean:
	rm -f generated/new_douyin/new_douyin.pb.go

help:
	@echo "Usage:"
	@echo "  make proto - Generate Go protobuf code"
	@echo "  make test  - Run Go tests"
	@echo "  make tidy  - Tidy Go dependencies"
	@echo "  make clean - Remove generated protobuf Go file"
