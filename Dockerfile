FROM golang:1.19 AS builder

COPY . /src
WORKDIR /src

RUN mkdir -p bin/ && go build -o ./bin/ ./...

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app

WORKDIR /app

ENTRYPOINT ["./app"]
