# twinte-back
[Twin:te](https://www.twinte.net/)のバックエンドです。

## Get Started
```sh
# prepare settings about environment variables
cp .env.example .env

# start docker containers
make up

# move into the container
make bash-app

# db migration
make migrate-up db_url=${DB_URL}
make migrate-up db_url=${TEST_DB_URL}

# start server
go run . serve
```

## Migration Tool
[golang-migrate](https://github.com/golang-migrate/migrate)というツールを使用しています。

Makefileで定義している主なコマンドは次の通りです。詳しくは`migrate --help`からご確認下さい。
```sh
# Apply all up migrations
make migrate-up db_url=${DB_URL}

# Apply all down migrations
make migrate-down db_url=${DB_URL}

# Set version but don't run migration
make migrate-force db_url=${DB_URL} version=1
```

### Installation
c.f. [CLI README](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

Goを用いてインスールを行う場合は、次のコマンドを実行して下さい。
```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

ビルド済みのバイナリからインストールを行う場合は、バージョン、OS、CPUアーキテクチャによってダウンロードURLが異なるため、[Releases](https://github.com/golang-migrate/migrate/releases)を参考にして適当なURLを指定して下さい。

バージョンが4.15.2、OSがLinux、CPUアーキテクチャがarm64のケースでのインストールは次のコマンドで実行できます。

```sh
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-arm64.tar.gz | tar xvz
mv ./migrate /usr/local/bin/
```
