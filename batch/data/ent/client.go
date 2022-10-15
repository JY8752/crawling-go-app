package data

import (
	"JY8752/crawling_app_batch/ent"
	"JY8752/crawling_app_batch/ent/migrate"
	"context"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewEntClient(ctx context.Context, connectionString string) *ent.Client {
	client, err := ent.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v\n", err)
	}

	//マイグレーションの実行
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v\n", err)
	}
	return client
}

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %v", rerr.Error())
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %v", err.Error())
	}
	return nil
}
