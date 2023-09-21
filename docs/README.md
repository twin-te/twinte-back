# Twinte Backend

## 開発方法
```
# docker containerを起動する
make up

# docker container内に入る
make bash

# DBのマイグレーションを行う
make migrate-up

# サーバーを起動する
make serve
```

## DB
[sqlboiler](https://github.com/volatiletech/sqlboiler)というツールを使用しています。ORMのようなものです。こちらのツールは事前に型情報を取得するためgormよりもパフォーマンスが良いらしいです。
ただ、batch insertやネストされたトランザクションがサポートされていないため、若干使いづらいです。

`make sqlboiler`を実行することでDBのテーブル構造からGoのコードを生成することができます。

マイグレーションには[migrate](https://github.com/golang-migrate/migrate)というツールを使用しています。


## Environment Variables
.envファイルを参照して下さい。

## API
[connect](https://connectrpc.com/docs/introduction)を使用しています。

1. `twinte-proto`配下にAPIの定義を記述する
2. `make buf-gen`を実行することでprotocol bufftersからGoのコードを生成する

## idtype
コード生成のためPythonが必要になります。Python 3.9では動きました。

```
# codegen/idtype/uuid_definition.txtにid名を定義する

# idtypeを生成する
python codegen/idtype/generate.py
```

## Software Architecture
<img width="758" alt="twinte_arch" src="https://github.com/twin-te/twinte-back/assets/68944024/7cf0cdd7-222d-489c-89a8-89c4514a29f7">

DDD・Clean Architecture・Modular Monolithを参考にして作成しました。

## OAuth2
OAuth2に関する設定は環境変数から行なっています。詳細は`.env`ファイルを参照して下さい。
