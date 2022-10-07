package main

import (
	"JY8752/crawling_app_batch/crawling"
	"JY8752/crawling_app_batch/utils"
	"fmt"
)

func main() {
	var (
		fu []string
		nu []string
	)
	urls, err := crawling.Crawling("http://example.com")
	if err != nil {
		fmt.Printf("failed crawling %v error %v\n", "http://example.com", err.Error())
	}

	nu = urls

	for {
		if len(nu) == 0 {
			break
		}

		for _, url := range nu {
			urls, err := crawling.Crawling(url)
			if err != nil {
				fmt.Printf("failed crawling %v error %v\n", url, err.Error())
				fu = append(fu, url)
				nu = utils.Remove(nu, url)
				continue
			}
			for _, u := range urls {
				if !utils.Contains(fu, u) {
					//未クローリング
					nu = append(nu, u)
				}
			}
		}
	}
}
