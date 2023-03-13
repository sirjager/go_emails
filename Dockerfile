FROM golang:1.19.7-alpine3.17 as builder
WORKDIR /app

RUN apk update
RUN apk add libc6-compat git build-base

COPY go.mod go.sum ./

RUN go mod download

COPY ./cfg ./cfg
COPY ./cmd ./cmd
COPY ./docs ./docs
COPY ./pkg ./pkg
COPY example.env .

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

RUN go build -o main ./cmd/main.go

FROM alpine:latest as runner
WORKDIR /app

RUN apk add libc6-compat

ENV GRPC_PORT=8030
ENV HTTP_PORT=8031

ENV TOKENS_GRPC_ADDR=localhost:8020

ENV SMTP_HOST=smtp.gmail.com
ENV SMTP_PORT=587
ENV SMTP_EMAIL=email@gmail.com
ENV SMTP_PASS=password

ENV API_HEADER=api-secret
ENV API_SECRET=emails

EXPOSE 8030
EXPOSE 8031

COPY --from=builder /app/main .
COPY --from=builder /app/example.env .

ENTRYPOINT [ "/app/main" ]
