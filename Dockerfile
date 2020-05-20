FROM golang:1.14-buster as goBuilder
WORKDIR /project
COPY . .
RUN rm -rf .env
RUN make

FROM debian:buster
WORKDIR /project
COPY --from=goBuilder /project/var/quiet-hacker-news /usr/local/bin/
COPY --from=goBuilder /project/templates/ ./templates
COPY --from=goBuilder /project/static/ ./static
RUN apt-get update
RUN apt-get install -y ca-certificates
CMD ["quiet-hacker-news"]
EXPOSE 9090
