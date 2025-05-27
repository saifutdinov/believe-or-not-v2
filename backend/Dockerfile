FROM golang:1.24-alpine

RUN apk add --no-cache \
    bash \
    curl \
    nano \
    htop \
    netcat-openbsd \
    redis \
    mc \
    less \
    procps \
    iputils \
    gcc \
    musl-dev \
    libc-dev


WORKDIR /app

COPY go.mod ./
RUN go mod download

RUN go install github.com/air-verse/air@latest

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.6

RUN golangci-lint --version