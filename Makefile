.PHONY: gen
gen:pbgen-go

.PHONY: pbgen-go
pbgen-publisher:
	protoc --proto_path=protobuf --go_out=go/pbdef --go_opt=paths=source_relative protobuf/definition.proto
