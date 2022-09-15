FROM golang:1.19-buster as goBuilder
WORKDIR /staging
COPY . .
RUN make full

FROM debian:buster
WORKDIR /app
RUN apt-get update
RUN apt-get install -y ca-certificates
COPY --from=goBuilder /staging/var/build ./build
CMD ["./build"]
EXPOSE 2222
