FROM golang:1.27.3@sha256:a386c44c1c90025ad5bc66feea3fef23052bef3bace3609f7e3e10e1f9314f45 AS build
WORKDIR /src

COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab

WORKDIR /app
USER 65532:65532

COPY --from=build /src/app .

CMD ["./app"]
