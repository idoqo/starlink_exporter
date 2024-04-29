.PHONY: test format build genproto
SHELL := /bin/bash
REQ_GEN_DEPS := grpcurl protoc protoc-gen-go protoc-gen-go-grpc
GEN_PROTO := spacex/api/common/status/status.proto spacex/api/device/command.proto spacex/api/device/common.proto spacex/api/device/device.proto spacex/api/device/dish_config.proto spacex/api/device/dish.proto spacex/api/device/services/unlock/service.proto spacex/api/device/transceiver.proto spacex/api/device/wifi_config.proto spacex/api/device/wifi_util.proto spacex/api/device/wifi.proto spacex/api/satellites/network/ut_disablement_codes.proto spacex/api/telemetron/public/common/time.proto
GEN_MODEL_DIR := pkg/spacex.com/api


test:
	go test -v -coverprofile=cover.out.tmp -coverpkg=./... ./...

format:	# Formats Go files
	goimports -local "github.com/idoqo/starlink_exporter/" -l -w ./
	for f in $(git status -s | cut -f3 -d ' '); do go fmt "$f"; done
	gci write --section Standard --section Default --section "Prefix(github.com/idoqo/starlink_exporter)" $(shell find . -type f -name "*.go")

build: # Build exporter binary
	go build -o ./dist/starlink_exporter ./cmd/starlink_exporter/main.go

genproto: # Generate proto files (thanks to github.com/b0ch3nski/go-starlink)
	$(foreach bin,$(REQ_GEN_DEPS),$(if $(shell which $(bin)),,$(error "Please install '$(bin)'")))

	grpcurl -plaintext -protoset-out "dist/protoset" "192.168.100.1:9200" describe SpaceX.API.Device.Device
	protoc --descriptor_set_in="dist/protoset" --go_out="$(GEN_MODEL_DIR)" --go-grpc_out="$(GEN_MODEL_DIR)" \
	--go_opt="module=spacex.com/api" --go-grpc_opt="module=spacex.com/api" \
	--go_opt="Mspacex/api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock" \
	--go-grpc_opt="Mspacex/api/device/services/unlock/service.proto=spacex.com/api/device/services/unlock" \
	--go_opt="Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites" \
	--go-grpc_opt="Mspacex/api/satellites/network/ut_disablement_codes.proto=spacex.com/api/satellites" \
	--go_opt="Mspacex/api/telemetron/public/common/time.proto=spacex.com/api/telemetron" \
	--go-grpc_opt="Mspacex/api/telemetron/public/common/time.proto=spacex.com/api/telemetron" \
	$(GEN_PROTO)


	find ./$(GEN_MODEL_DIR) -type f -name "*.go" -exec sed -i.old "s|spacex.com/api|$$(awk '/module/{print $$2}' go.mod)/$(GEN_MODEL_DIR)|g" {} \; # Provide sed with suffix (.old instead of just plain -i) to work on MacOS.
	find ./$(GEN_MODEL_DIR) -type f -name "*.old" -exec rm {} \;
