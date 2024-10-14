FROM golang:1.21.6 AS builder


ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go mod tidy

COPY . .


RUN go build -o server cmd/server/main.go



FROM gcr.io/distroless/base-debian11

COPY --from=builder /app/server .


CMD ["./server"]