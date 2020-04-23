.PHONY:all
all:pointerpw

.PHONY:pointerpw
pointerpw: mac linux

.PHONY:proto
proto:
	protoc -I./pb pb/coredns.proto --go_out=plugins=grpc:./pb
	protoc -I./pb pb/dnsregistry.proto --go_out=plugins=grpc:./pb
	protoc -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.4/third_party/googleapis -I./pb pb/ssrf.proto --go_out=plugins=grpc:./pb
	protoc -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.4/third_party/googleapis -I./pb pb/ssrf.proto --grpc-gateway_out=:./pb

.PHONY:mac
mac: proto
	GOOS=darwin GOARCH=amd64 go build -o build/mac cli/main.go


.PHONY:linux
linux: proto
	rm -r build/linux || 0
	mkdir -p build/linux
	GOOS=linux GOARCH=amd64 go build -o build/linux/pointerpw cli/main.go
	cp pkg/linux/install.sh build/linux/
	cp pkg/linux/pointerpw.service build/linux/
	cp pkg/linux/api.pointer.pw build/linux/
	cp pkg/linux/update.sh build/linux/
	cd build/linux && zip -r linux.amd64.zip ./**
