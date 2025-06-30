# redis

gofrでRedisにアクセする場合のサンプル。

## 実行方法

起動方法:

```shell
make up
```

停止方法:

```shell
make down
```

## 試験方法

```shell
# Redisのコンテナを起動
make redis-up
# ユニットテストを実行
make test
# Redisのコンテナを停止
make redis-down
```
