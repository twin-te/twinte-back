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

# install migration tool
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-arm64.tar.gz | tar xvz
mv ./migrate /usr/local/bin/

# db migration
make migrate-up db_url=${DB_URL}
make migrate-up db_url=${TEST_DB_URL}

# start server
go run . serve
```

## Migration Tool
[golang-migrate](https://github.com/golang-migrate/migrate)というツールを使用しています。

バージョン、OS、CPUアーキテクチャによってダウンロードURLが異なるため、下記のリンクを参考にして適当なURLを指定して下さい。

- [golang-migrate - README](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)
- [golang-migrate - Releases](https://github.com/golang-migrate/migrate/releases)


バージョンが4.15.2、OSがLinux、CPUアーキテクチャがarm64のケースでのURLは次のようになります。

```
https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-arm64.tar.gz
```

主なコマンドは次の通りです。詳しくは`migrate --help`からご確認下さい。
```sh
# Apply all up migrations
make migrate-up db_url=${DB_URL}

# Apply all down migrations
make migrate-down db_url=${DB_URL}

# Set version but don't run migration
make migrate-force db_url=${DB_URL} version=1
```
