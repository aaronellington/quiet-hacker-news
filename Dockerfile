FROM golang:1.16-buster as goBuilder
WORKDIR /build-staging
COPY . .
RUN make clean lint test build

FROM debian:buster
RUN apt-get update
RUN apt-get install -y ca-certificates
WORKDIR /app
COPY --from=goBuilder /build-staging/var/quiet-hacker-news ./quiet-hacker-news
COPY --from=goBuilder /build-staging/assets/ ./assets
CMD ["./quiet-hacker-news"]
EXPOSE 8000
