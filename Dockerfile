FROM golang:1.27.4@sha256:fe517acaa580cc7915457080f850c37e618b3c1e750f98b58fc6429b605147a1 AS build
WORKDIR /src

COPY . .
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab

WORKDIR /app
USER 65532:65532

COPY --from=build /src/app .

CMD ["./app"]
