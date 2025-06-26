#build stage
FROM golang:1.23-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN apk add tar
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.1/migrate.linux-arm64.tar.gz | tar xvz

#run stage
FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/APP.env .
COPY --from=builder /app/migrate .
COPY --from=builder /app/db/migration ./db/migration

RUN mv migrate /usr/bin/migrate
RUN which migrate
EXPOSE 8080
CMD [ "/app/main"]