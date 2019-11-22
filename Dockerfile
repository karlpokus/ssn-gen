# 1. build binary

FROM golang:alpine AS builder

# install git to fetch deps
RUN apk update && apk add --no-cache git

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download # deps will be cached

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ssns ./cmd/main

# 2. copy only binary to new image

FROM scratch

WORKDIR /src

COPY --from=builder /src/ssns ./ssns

EXPOSE 9012

ENTRYPOINT ["/src/ssns"]
