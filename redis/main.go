package main

import (
	"time"

	"gofr.dev/pkg/gofr"
)

const redisExpiryTime = 5

func main() {
	// Create a new application
	app := gofr.New()

	// Add routes for Redis operations
	app.GET("/redis/{key}", RedisGetHandler)
	app.POST("/redis", RedisSetHandler)
	app.GET("/redis-pipeline", RedisPipelineHandler)

	// Run the application
	app.Run()
}

// RedisSetHandler はリクエストからキーと値を取得し、Redis に保存するハンドラーです。
// データは `redisExpiryTime` 分間有効です。
// リクエストのバインドや Redis 操作中にエラーがあれば、それを返します。
func RedisSetHandler(c *gofr.Context) (any, error) {
	input := make(map[string]string)
	// Requestの内容をinputにバインドする
	if err := c.Request.Bind(&input); err != nil {
		return nil, err
	}

	for key, value := range input {
		// Redisにkey,valueのペアで値を追加
		err := c.Redis.Set(c, key, value, redisExpiryTime*time.Minute).Err()
		if err != nil {
			return nil, err
		}
	}

	return "Successful", nil
}

// RedisGetHandler は指定されたキーに基づいて Redis から値を取得し、結果をマップ形式で返すハンドラーです。
// 存在しないキーや Redis 操作中にエラーが発生した場合はエラーを返します。
func RedisGetHandler(c *gofr.Context) (any, error) {
	key := c.PathParam("key")
	// Redisから値を取得
	value, err := c.Redis.Get(c, key).Result()
	if err != nil {
		return nil, err
	}
	// 戻り値をマップにする
	resp := make(map[string]string)
	resp[key] = value

	return resp, nil
}

// RedisPipelineHandler は複数の Redis コマンドをパイプライン内で実行し、結果を返すハンドラーです。
// エラーが発生した場合は即時に処理を終了し、エラーを返します。
func RedisPipelineHandler(c *gofr.Context) (any, error) {
	// パイプラインを作成
	pipe := c.Redis.Pipeline()

	// パイプラインに複数のコマンドを追加
	pipe.Set(c, "testKey1", "testValue1", redisExpiryTime*time.Minute)
	pipe.Get(c, "testKey1")

	// パイプラインを実行
	cmds, err := pipe.Exec(c)
	if err != nil {
		return nil, err
	}

	// パイプライン内の各コマンドの結果を処理するか返却する（実装は省略）
	return cmds, nil
}
