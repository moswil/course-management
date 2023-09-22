PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
PROTO_DIR = internal/app/grpc/proto

out:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. \
  --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. ${PROTO_DIR}/*.proto
