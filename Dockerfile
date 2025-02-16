FROM golang:1.23 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY cmd /app/cmd
RUN cd cmd/server && go build -o server main.go

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/cmd/server/server /app/server/

WORKDIR /app

EXPOSE 8080
CMD ["/app/server/server"]
