FROM golang:latest AS builder

WORKDIR /build
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server

FROM alpine:latest

WORKDIR /server
COPY --from=builder /build/server .

CMD ["./server"]
