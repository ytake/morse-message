.PHONY: gen
gen:pbgen-publisher pbgen-consumer

.PHONY: pbgen-publisher
pbgen-publisher:
	protoc --proto_path=protobuf --go_out=publisher/pbdef --go_opt=paths=source_relative protobuf/definition.proto

.PHONY: pbgen-consumer
pbgen-consumer:
	protoc --proto_path=protobuf --go_out=consumer/pbdef --go_opt=paths=source_relative protobuf/definition.proto
