FROM golang:1.19-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

ENV CGO_ENABLED=0

RUN go build -o broker_app .

RUN chmod +x /app/broker_app

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/broker_app /app

CMD ["/app/broker_app"]
