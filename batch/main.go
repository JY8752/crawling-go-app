package main

import (
	"JY8752/crawling_app_batch/crawling"
	"JY8752/crawling_app_batch/utils"
	"fmt"
	"time"
)

func main() {
	targetURL := "https://lycoris-recoil.com/"
	fmt.Printf("start crwling. %v\n", targetURL)
	var (
		crawledUrls []string //クローリング済みのURL
		nextUrls    []string //クローリング待機URL
		err         error
		counter     int
	)

	//エントリーポイントクローリング
	nextUrls, err = crawling.Crawling(targetURL)
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
			if utils.Contains(crawledUrls, url) {
				nextUrls = utils.Remove(nextUrls, url)
				continue
			}

			//ページ巡回する
			fmt.Printf("crawling page url: %v\n", url)
			urls, err := crawling.Crawling(url)
			fmt.Printf("extract link urls urls: %v\n", urls)

			crawledUrls = append(crawledUrls, url)
			nextUrls = utils.Remove(nextUrls, url)

			//クローリングに失敗したら巡回ずみにして次へ
			if err != nil {
				fmt.Printf("failed crawling %v error %v\n", url, err.Error())
				continue
			}

			//未クローリングURLであれば巡回リストに追加する
			for _, u := range urls {
				if !utils.Contains(crawledUrls, u) {
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
