package main

import (
	"JY8752/crawling_app_batch/array"
	"JY8752/crawling_app_batch/crawling"
	data "JY8752/crawling_app_batch/data/ent"
	_ "JY8752/crawling_app_batch/env"
	"JY8752/crawling_app_batch/http"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	//DB接続
	ctx := context.Background()
	cs := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DOMAIN"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)
	data.NewEntClient(ctx, cs)

	targetURL := "https://lycoris-recoil.com/"
	fmt.Printf("start crwling. %v\n", targetURL)
	var (
		crawledUrls []string //クローリング済みのURL
		nextUrls    []string //クローリング待機URL
		err         error
		counter     int
	)

	c := http.NewClient()
	cr := crawling.NewClawler(c)

	//エントリーポイントクローリング
	nextUrls, err = cr.Crawling(targetURL)
	if err != nil {
		fmt.Printf("failed crawling %v error %v\n", targetURL, err.Error())
	}

	//全てのリンクを巡回する or 5ページ巡回するまでループ
loop:
	for {
		counter++
		fmt.Printf("-------------------- loop %d -------------------\n", counter)
		fmt.Printf("crawledUrls: %v nextUrls: %v\n", crawledUrls, nextUrls)
		if len(nextUrls) == 0 {
			break
		}

		for _, url := range nextUrls {
			//巡回ページが5ページになったら終了
			if len(crawledUrls) > 5 {
				break loop
			}

			//ブロックされないように少し待つ
			time.Sleep(time.Second * 3)
			//クローリング済みであれば巡回リストから除去して次へ
			if array.Contains(crawledUrls, url) {
				nextUrls = array.Remove(nextUrls, url)
				continue
			}

			//ページ巡回する
			fmt.Printf("crawling page url: %v\n", url)
			urls, err := cr.Crawling(url)
			fmt.Printf("extract link urls urls: %v\n", urls)

			crawledUrls = append(crawledUrls, url)
			nextUrls = array.Remove(nextUrls, url)

			//クローリングに失敗したら巡回ずみにして次へ
			if err != nil {
				fmt.Printf("failed crawling %v error %v\n", url, err.Error())
				continue
			}

			//未クローリングURLであれば巡回リストに追加する
			for _, u := range urls {
				if !array.Contains(crawledUrls, u) {
					//未クローリング
					nextUrls = append(nextUrls, u)
				} else {
					//クローリング済み
					fmt.Printf("already crawling url %v\n", u)
				}
			}
		}
	}
	fmt.Println("end crwling.")
}
