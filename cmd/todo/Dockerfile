FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build ./cmd/todo

FROM debian

WORKDIR /app
COPY --from=builder /build/todo ./todo
CMD ["/app/todo"]