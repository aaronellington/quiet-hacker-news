## Build the Go package
FROM golang:1.21-bullseye as goBuilder
WORKDIR /workspace/
COPY . .
RUN make clean build-go

FROM debian:bullseye
WORKDIR /workspace/
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=goBuilder /workspace/var/build ./build
CMD ["./build"]
EXPOSE 2222
