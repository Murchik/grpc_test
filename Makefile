PROTO_DIR_NAME := databus
PROTO_FILE_NAME := databus

all:

run:

build: proto

proto: $(PROTO_DIR_NAME)
	protoc --go_out=$(PROTO_DIR_NAME) --go_opt=paths=source_relative \
    --go-grpc_out=$(PROTO_DIR_NAME) --go-grpc_opt=paths=source_relative \
    $(PROTO_FILE_NAME).proto

databus:
	mkdir $(PROTO_DIR_NAME)

clean:
	rm -rf $(PROTO_DIR_NAME)
