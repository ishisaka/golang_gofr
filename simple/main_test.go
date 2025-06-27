package main

import (
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gofr.dev/pkg/gofr/testutil"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("GOFR_TELEMETRY", "false")
	m.Run()
}

func TestGreetEndpoint(t *testing.T) {
	// NewServerConfigs は、テスト用のサーバー設定を構成し、ServiceConfigs 構造体を返します。
	// この関数は、HTTP、Metrics、および gRPC サービス用に利用可能なポートを動的に割り当て、
	// それらの環境変数を設定し、構成された値を含む構造体を返します。
	configs := testutil.NewServerConfigs(t)

	// HTTPクライアントを作成
	c := &http.Client{}

	// サーバーを起動
	go main()
	time.Sleep(100 * time.Millisecond)

	testCases := []struct {
		desc        string
		method      string
		path        string
		statusCode  int
		expectedRes string
	}{
		{
			desc:        "greet endpoint",
			method:      http.MethodGet,
			path:        "/greet",
			statusCode:  http.StatusOK,
			expectedRes: "{\"data\":\"Hello World!\"}\n",
		},
		{
			desc:        "incorrect method",
			method:      http.MethodPost,
			path:        "/greet",
			statusCode:  http.StatusNotFound,
			expectedRes: "{\"error\":{\"message\":\"route not registered\"}}\n",
		},
		{
			desc:        "not found endpoint",
			method:      http.MethodGet,
			path:        "/not-found",
			statusCode:  http.StatusNotFound,
			expectedRes: "{\"error\":{\"message\":\"route not registered\"}}\n",
		},
	}

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			req, _ := http.NewRequest(tc.method, configs.HTTPHost+tc.path, nil)
			resp, err := c.Do(req)

			require.NoError(t, err, "TEST[%d], Failed.\n%s", i, tc.desc)

			bodyBytes, err := io.ReadAll(resp.Body)

			require.NoError(t, err, "TEST[%d], Failed.\n%s", i, tc.desc)
			assert.Equal(t, tc.expectedRes, string(bodyBytes), "TEST[%d], Failed.\n%s", i, tc.desc)
			assert.Equal(t, tc.statusCode, resp.StatusCode, "TEST[%d], Failed.\n%s", i, tc.desc)

			_ = resp.Body.Close()
		})
	}
}
