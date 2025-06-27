# シンプルなサービス

[Quick Start Guide \| Hello Server](https://gofr.dev/docs/quick-start/introduction)のコード。

## 実行方法

```bash
go mod tidy

```

```bash
go run main.go
```

## 設定切り替え

[Quick Start Guide \| Configuration](https://gofr.dev/docs/quick-start/configuration)

gofrは`config/.env`にて各種設定を行う。

以下は設定の例:

```aiignore
# configs/.env

APP_NAME=test-service
HTTP_PORT=9000
```

また、この設定は`APP_ENV`環境変数で変更できます。たとえば`.dev.env`という設定ファイルを作成し、これを実行時に使用したい場合には、以下のように実行する。

```shell
APP_ENV=dev go run main.go
```
設定項目の詳細は以下を参照:

[References \| Configs](https://gofr.dev/docs/references/configs)

