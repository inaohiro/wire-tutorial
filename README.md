# wire-tutorial

## Requirements

- wire : https://github.com/google/wire

### ディレクトリ構造

Clean Architecture を参考にしましたが、まだまだ理解が浅いので、usecase, interface 層に対応するディレクトリはありません

```
.
├─.idea
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
