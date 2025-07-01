# mysql

MySQLアクセスのサンプル。

## 実行手順

MySQLの起動

```shell
make mysqlup
```

データベースの初期化

```shell
nake migrate
```

サーバーの起動

```shell
make run
```

サーバーアクセス例

```shell
# 追加
curl -v -X POST localhost:9000/customer/bar

# 取得
curl -v -X GET localhost:9000/customer/1
curl -v -X GET localhost:9000/customer
```

後始末

```shell
make mysqldown
```