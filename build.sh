#!/bin/bash
BUILD_SOURCE="./main.go"
BUILD_DIRECTORY="./bin"
BUILD_OUTPUT="./bin/app"
GO_DEPENDENCY_FILE="go.mod"

# Make sure source code is available.
f=$BUILD_SOURCE && if test -f ${f}; then echo "found build source ${f}"; else echo "can't find build source: ${f} exiting" && exit 1; fi
# Check dependency management.
f=$GO_DEPENDENCY_FILE && if test -f ${f}; then echo "found ${f} running go mod tidy" && go mod tidy; else echo "can't find ${f} running go mod init and go mod tidy" && go mod init && go mod tidy; fi
# Clear out old builds if exist
d=$BUILD_DIRECTORY && if test -d ${d}; then echo "found build directory ${d} clearing folder" && rm ${d}/*; else mkdir ${d}; fi
# Build application
go build -o $BUILD_OUTPUT $BUILD_SOURCE
# Verify build output success
f=$BUILD_OUTPUT && if test -f ${f}; then echo "found build output success ${f}"; else echo "can't find build output ${f}, something went wrong" && exit 1; fi