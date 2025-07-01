// DBアクセスサンプル
// https://gofr.dev/docs/quick-start/connecting-mysql
// GoFrは、サポートされているすべてのSQLデータベースに対して、データとの
// 相互作用のための一貫したAPIを提供します。
//
// 次に、以下の例では、POST /customerを使用して顧客データを保存し、その後
// GET /customerを使用して同じデータを取得します。顧客データはIDと名前で保
// 存されます。
//
// SQLデータストアからデータを読み書きするためのコードを追加した後、main.goは
// 以下のように更新されます。

package main

import (
	"errors"

	"gofr.dev/pkg/gofr"
)

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// initialize gofr object
	app := gofr.New()

	app.POST("/customer/{name}", func(ctx *gofr.Context) (any, error) {
		name := ctx.PathParam("name")

		// SQLを使ってDBにデータを挿入する
		_, err := ctx.SQL.ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)

		return nil, err
	})

	app.GET("/customer", func(ctx *gofr.Context) (any, error) {
		var customers []Customer

		// SQLを使ってDBからデータを取得する
		rows, err := ctx.SQL.QueryContext(ctx, "SELECT * FROM customers")
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var customer Customer
			if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
				return nil, err
			}

			customers = append(customers, customer)
		}

		// return the customer
		return customers, nil
	})

	app.GET("/customer/{id}", func(ctx *gofr.Context) (any, error) {
		var customers []Customer

		// SQLを使ってDBからデータを取得する
		rows, err := ctx.SQL.QueryContext(ctx, "SELECT * FROM customers WHERE id = ?", ctx.PathParam("id"))
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var customer Customer
			if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
				return nil, err
			}

			customers = append(customers, customer)
		}
		// 2件以上取得できた場合はエラー
		if len(customers) > 1 {
			return nil, errors.New("multiple customers found")
		}
		return customers[0], nil
	})

	app.Run()
}
