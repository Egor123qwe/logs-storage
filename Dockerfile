FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -C cmd/logsStorage/ -o /logsstorage


FROM alpine:3.19

WORKDIR /app

COPY --from=builder /logsstorage /app/logsstorage
COPY config/        /app/config/

EXPOSE 8081

ENTRYPOINT [ "/app/logsstorage" ]
