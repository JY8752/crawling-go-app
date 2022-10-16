package main

import (
	"context"
	"log"
	"net/http"

	"JY8752/crawling_app_rest/presentation"
)

func main() {
	ctx := context.Background()

	//route登録
	presentation.Route(ctx)

	//server起動
	log.Fatal(http.ListenAndServe(":8080", nil))
}
