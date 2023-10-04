PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
PROTO_DIR = internal/app/grpc/proto

gen_grpc_protoc:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. \
  --go-grpc_opt=module=${PACKAGE} --go-grpc_out=. ${PROTO_DIR}/*.proto

gen_event_pb:
	protoc -I internal/core/domain/event/proto --go_opt=module=${PACKAGE} \
	--go_out=. internal/core/domain/event/proto/*.proto