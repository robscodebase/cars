FROM golang:1.21 as build

WORKDIR /go/src/cars
COPY . .

RUN echo "STARTING BUILD" && echo "Verifying env file, go mod, and bin folder." \
    && f="./.env" && if test -f ${f}; then echo "found ${f}"; else echo "couldn't find ${f}"; fi \
    && f="./go.mod" && if test -f ${f}; then echo "found ${f}"; else go mod init; fi \
    && d="./bin" && if test -d ${d}; then echo "found ${d}" && rm -r ${d}; fi \
    && BUILD_DIR="./build-files" \
    && d="${BUILD_DIR}/bin" && if test -d ${d}; then echo "found ${d}" && mkdir -p ${d}; else mkdir -p ${d}; fi \
    && echo "Copying .env file to build-files folder." \
    && cp ./.env ${BUILD_DIR} \
    && FINAL_BUILD="${BUILD_DIR}/bin/app" \
    && echo "Verifying main.go location." \
    && f="./main.go" && if test -f ${f}; then echo "found ${f}" && go build -o ${FINAL_BUILD} ${f}; fi\
    && echo "Build Finished" \
    && echo "Verifying build" \
    && f="${FINAL_BUILD}" && if test -f ${f}; then echo "found ${f}"; else echo "couldn't find ${f}"; fi


FROM debian:latest AS deploy
COPY --from=build /go/src/cars/build-files /go/

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

ENTRYPOINT /go/bin/app
WORKDIR /go

EXPOSE 8080
