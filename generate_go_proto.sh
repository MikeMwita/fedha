#!/bin/bash

# Generate Go Protobuf and gRPC code for the app-auth module
protoc --proto_path=protos \
       --go_out=services/app-auth/pb \
       --go-grpc_out=services/app-auth/pb \
       $(find protos/app-auth -name '*.proto')

# Generate Go Protobuf and gRPC code for the app-db module
protoc --proto_path=protos \
       --go_out=services/app-db/pb \
       --go-grpc_out=services/app-db/pb \
       $(find protos/app-db -name '*.proto')

# Generate Go Protobuf and gRPC code for the app-expense module
protoc --proto_path=protos \
       --go_out=services/app-expense/pb \
       --go-grpc_out=services/app-expense/pb \
       $(find protos/app-expense -name '*.proto')
