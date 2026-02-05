FROM golang:1.25.5@sha256:8bbd14091f2c61916134fa6aeb8f76b18693fcb29a39ec6d8be9242c0a7e9260 AS build
WORKDIR /src

COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest@sha256:25109184c71bdad752c8312a8623239686a9a2071e8825f20acb8f2198c3f659

RUN addgroup -g 1000 app && \
    adduser -u 1000 -h /app -G app -S app
WORKDIR /app
USER app

COPY --from=build /src/app .

CMD ["./app"] 
