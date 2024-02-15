FROM --platform=linux/amd64 golang:1.21-bookworm as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -v -o server ./cmd/main.go

FROM --platform=linux/amd64 debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
COPY --from=builder /app/config/docker-config.yml /app/config/docker-config.yml
COPY --from=builder /app/config/countries.yml /app/config/countries.yml

ENV APP_ENV docker
ENV PORT ${PORT}

RUN chmod +x /app/server

CMD ["/app/server"]