package main

import (
	"JY8752/crawling_app_batch/array"
	"JY8752/crawling_app_batch/crawling"
	crawledUrl "JY8752/crawling_app_batch/data/crawledurl"
	db "JY8752/crawling_app_batch/data/db"
	data "JY8752/crawling_app_batch/data/ent"
	linkUrl "JY8752/crawling_app_batch/data/linkurl"
	"JY8752/crawling_app_batch/http"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type NextUrl struct {
	url     string
	referer string
}

func main() {
	godotenv.Load()

	//DB接続
	ctx := context.Background()
	cs := db.GetConnectionStr()
	client := data.NewEntClient(ctx, cs)

	cu := crawledUrl.NewCrawledUrl(client)
	lu := linkUrl.NewLinkUrl(client)

	targetURL := strings.TrimSuffix("https://lycoris-recoil.com/", "/")
	fmt.Printf("start crwling. %v\n", targetURL)
	var (
		crawledUrls []string  //クローリング済みのURL
		nextUrls    []NextUrl //クローリング待機URL
		err         error
		counter     int
	)

	c := http.NewClient()
	cr := crawling.NewClawler(c)

	//エントリーポイントクローリング
	result, err := cr.Crawling(targetURL)
	if err != nil {
		log.Fatalf("failed crawling %v error %v\n", targetURL, err.Error())
	}

	for _, u := range result.LinkUrls {
		nextUrls = append(nextUrls, NextUrl{url: u, referer: targetURL})
	}

	//起点URLを登録
	s, err := cu.Save(ctx, targetURL, nil)
	if err != nil {
		log.Fatalf("failed save crawledurl. url: %v error %v\n", targetURL, err.Error())
	}

	//全てのリンクを巡回する or 5ページ巡回するまでループ
loop:
	for {
		counter++
		fmt.Printf("crawledUrls: %v nextUrls: %v\n", crawledUrls, nextUrls)
		if len(nextUrls) == 0 {
			break
		}

		for i, url := range nextUrls {
			fmt.Printf("---------------- count: %v url: %v crawled: %v -------------------\n", i, url, len(crawledUrls))
			//巡回ページが5ページになったら終了
			if len(crawledUrls) > 4 {
				break loop
			}

			//ブロックされないように少し待つ
			time.Sleep(time.Second * 3)
			//クローリング済みであれば巡回リストから除去して次へ
			if array.Contains(crawledUrls, url.url) {
				nextUrls = remove(nextUrls, url)
				continue
			}

			//ページ巡回する
			fmt.Printf("crawling page url: %v\n", url)
			result, err := cr.Crawling(url.url)

			crawledUrls = append(crawledUrls, url.url)
			nextUrls = remove(nextUrls, url)

			//クローリングに失敗したら巡回ずみにして次へ
			if err != nil {
				fmt.Printf("failed crawling %v error %v\n", url, err.Error())
				continue
			}

			fmt.Printf("extract link urls urls: %v\n", result.LinkUrls)

			//クローリング結果を登録
			lu.Save(ctx, url.url, url.referer, nil, s)

			//未クローリングURLであれば巡回リストに追加する
			for _, u := range result.LinkUrls {
				if !array.Contains(crawledUrls, u) {
					//未クローリング
					nextUrls = append(nextUrls, NextUrl{url: u, referer: url.url})
				} else {
					//クローリング済み
					fmt.Printf("already crawling url %v\n", u)
				}
			}
		}
	}
	fmt.Println("end crwling.")
}

func remove(urls []NextUrl, url NextUrl) []NextUrl {
	var arr []NextUrl
	for _, u := range urls {
		if u.url != url.url {
			arr = append(arr, u)
		}
	}
	return arr
}
