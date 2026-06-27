FROM golang:1.26-alpine3.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -ldflags="-s -w" .

FROM scratch
COPY --from=alpine:3.24 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/svc-discord /svc-discord

ENV GIN_MODE=release
CMD ["/svc-discord"]
