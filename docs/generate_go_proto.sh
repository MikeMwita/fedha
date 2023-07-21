#!/usr/bin/env bash


#!/bin/bash

DB_PROTO_DIR="./db"
EXPENSE_PROTO_DIR="./expense"
OUTPUT_DIR="./generated_rpc_code"

mkdir -p "$OUTPUT_DIR"

# Generate code for db directory
protoc --proto_path="$DB_PROTO_DIR" --go_out="$OUTPUT_DIR" --go-grpc_out="$OUTPUT_DIR" "$DB_PROTO_DIR"/*.proto

# Generate code for expense directory
protoc --proto_path="$EXPENSE_PROTO_DIR" --go_out="$OUTPUT_DIR" --go-grpc_out="$OUTPUT_DIR" "$EXPENSE_PROTO_DIR"/*.proto
