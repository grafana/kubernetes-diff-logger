FROM golang:1.27.1@sha256:3680233e3204827fbdc66088528ae6d4b3d034f51d03a99d454f6de034888244 AS build
WORKDIR /src

COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest

RUN addgroup -g 1000 app && \
    adduser -u 1000 -h /app -G app -S app
WORKDIR /app
USER app

COPY --from=build /src/app .

CMD ["./app"] 
