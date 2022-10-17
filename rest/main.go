package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	infrastructure "JY8752/crawling_app_rest/infrastructure/ent"
	"JY8752/crawling_app_rest/presentation"

	"github.com/joho/godotenv"
)

func main() {
	//DB接続
	godotenv.Load()
	ctx := context.Background()

	cs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%v",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DOMAIN"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
		"Asia%2FTokyo",
	)
	client := infrastructure.NewEntClient(ctx, cs)

	//route登録
	presentation.Route(ctx, client)

	//server起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
