FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM busybox

COPY --from=builder /app/main /main

EXPOSE 8081

CMD ["./main"]
