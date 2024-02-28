FROM golang:1.21 as dev

WORKDIR /usr/src/app

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
  && go install github.com/bufbuild/buf/cmd/buf@latest \
  && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest \
  && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
  && go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest \
  && go install github.com/spf13/cobra-cli@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

CMD [ "sleep", "infinity" ]
