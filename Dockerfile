FROM golang:1.25 AS builder

WORKDIR /app
COPY . .
RUN ./build.sh 

FROM alpine:latest


WORKDIR /root/
COPY --from=builder /app/bin/ .