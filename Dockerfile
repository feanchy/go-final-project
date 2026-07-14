FROM golang:1.26 AS builder

WORKDIR /app

COPY . .

RUN go build -o scheduler .


FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/scheduler .
COPY web ./web

EXPOSE 7540

CMD ["./scheduler"]