package main

import "gofr.dev/pkg/gofr"

func main() {
	// gofrサーバーの作成
	// gofr.New() が呼び出されると、フレームワークを初期化し、
	// 設定ファイルに基づいてロガー、メトリクス、データソース
	// などの初期化処理を実行します。
	app := gofr.New()

	// 特定のパスとハンドラー関数の接続を行う
	// GoFr コンテキスト ctx *gofr.Context は、リクエスト、レスポンス、
	// 依存関係をラップするクラスで、さまざまな機能を提供します。
	app.GET("/greet", func(ctx *gofr.Context) (any, error) {
		return "Hello World!", nil
	})

	// サーバーを起動する
	// app.Run() が呼び出されると、HTTP サーバーとミドルウェアの構成、初期化、
	// 実行を行います。ヘルスチェックエンドポイントのルーティング、メトリクス
	// サーバー、ファビコンなどの重要な機能を管理します。設定が無ければデフォルトの
	// ポート 8000 でサーバーを起動します。
	app.Run()
}
