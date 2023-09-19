# Development
FROM golang:1.20 as dev

WORKDIR /usr/src/app

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-arm64.tar.gz | tar xvz \
  && mv ./migrate /usr/local/bin/ \
  && go install github.com/cosmtrek/air@latest \
  && go install github.com/volatiletech/sqlboiler/v4@latest \
  && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest \
  && go install github.com/bufbuild/buf/cmd/buf@latest \
  && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest \
  && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest \
  && go install github.com/spf13/cobra-cli@latest \
  && go install golang.org/x/tools/gopls@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

CMD [ "sleep", "infinity" ]
