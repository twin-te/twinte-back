FROM golang:1.21-bullseye as builder

WORKDIR /app

COPY . .

RUN go mod download \
  && go mod verify \
  && go build -trimpath -ldflags "-w -s" -o app

FROM debian:bullseye-slim as deploy

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]
