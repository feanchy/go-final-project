FROM golang:1.26 AS builder

WORKDIR /app

COPY . .

RUN go build -o scheduler .


FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /app/scheduler .
COPY web ./web

ENV TODO_PORT=7540
ENV TODO_DBFILE=/app/scheduler.db
ENV TODO_PASSWORD=1234

EXPOSE 7540

CMD ["./scheduler"]