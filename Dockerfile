# syntax=docker/dockerfile:1
FROM golang:1.22.3 as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ../.. ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/server ./cmd/server


FROM gcr.io/distroless/static@sha256:41972110a1c1a5c0b6adb283e8aa092c43c31f7c5d79b8656fbffff2c3e61f05 as final
WORKDIR /app

EXPOSE 8081

COPY --from=builder /bin/server /server
CMD ["/server"]
