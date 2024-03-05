FROM golang:1.21-bullseye as builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy
RUN go build -trimpath -ldflags "-w -s" -o app

FROM debian:bullseye-slim

RUN apt-get update \
  && DEBIAN_FRONTEND=noninteractive apt-get install -yq ca-certificates openssl \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app", "serve"]
