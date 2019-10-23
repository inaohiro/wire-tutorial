# wire-tutorial

## Requirements

- wire : https://github.com/google/wire


## How to run

```
git clone https://github.com/inaohiro/wire-tutorial.git
cd wire-tutorial/app
wire
go build
./app
```

## ディレクトリ構造

Clean Architecture を参考にしましたが、まだまだ理解が浅いので、usecase, interface 層に対応するディレクトリはありません

```
.
└─app
    ├─config
    ├─domain           # domain 以下にロジックを置きます
    │  ├─application   #   アプリケーションのロジック
    │  └─domain        #   モデルおよびリポジトリのインターフェース
    │
    └─infra            # infra 以下にデータベースや REST API のルーティングなどを置きます
        ├─database     # ここでは MySQL に接続するものと Mock 用 の InMemoryDB を用意しています
        └─inmemorydb
```
