package presentation

import (
	"JY8752/crawling_app_rest/presentation/controller"
	"context"
)

func Route(ctx context.Context) {
	//crawling route
	controller.NewCrawlingController().Register()
}
