#!/bin/bash

echo 'Generating proto files...'
docker run --rm -v //$(pwd)/api:/$(pwd)/api -w /$(pwd)/api znly/protoc \
    --go_out=plugins=grpc:/$(pwd)/api \
    -I. person.proto

echo 'Proto files generated ㋡'