deps:
	@echo "--> Installing dependencies for Plasma Core..."
	@echo "--> Installing Go dependencies..."
	@go mod tidy

build-all: build-plasmad build-plasmacli build-harness

build-plasmad:
	go build -o ./target/plasmad ./cmd/plasmad/main.go

build-harness:
	go build -o ./target/plasma-harness ./cmd/harness/main.go

build-plasmacli:
	go build -o ./target/plasmacli ./cmd/plasmacli/main.go

build-plasmad-debug:
	go build -gcflags "all=-N -l" -o ./target/plasmad ./cmd/plasmad/main.go

install:
	go install ./cmd/plasmad

protogen:
	protoc -I pkg/rpc/proto pkg/rpc/proto/root.proto --go_out=plugins=grpc:pkg/rpc/pb

clean:
	rm -rf ./integration_tests/node_modules
	rm -rf ./pkg/eth/contracts/plasma_mvp.go
	rm -rf ./pkg/eth/contracts/PlasmaMVP.abi
	rm -rf ./target
	rm -rf ~/.plasma
	rm -rf .vendor-new

test-integration:
	npm --prefix ./integration_tests test

test-unit:
	go test -v ./...

.PONY: build deps test package
