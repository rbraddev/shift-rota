# syntax=docker/dockerfile:1

FROM golang:1.18 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags='-s' -o app ./cmd/api

FROM alpine:latest AS production

RUN mkdir /opt/app
WORKDIR /opt/app

COPY --from=builder /app/app /app/entrypoint.sh ./

ENTRYPOINT [ "/opt/app/entrypoint.sh" ]