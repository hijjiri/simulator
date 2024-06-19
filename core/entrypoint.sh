#!/bin/sh

# Substitute the PORT environment variable in the gRPC server command
if [ -z "$PORT" ]; then
  PORT=50051
fi

grpc-server -port $PORT
