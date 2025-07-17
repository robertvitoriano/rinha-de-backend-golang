FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -x -v -o rinha cmd/main.go

FROM debian:bookworm

WORKDIR /app

COPY --from=builder /app/rinha .

EXPOSE 4444

CMD ["./rinha"]
