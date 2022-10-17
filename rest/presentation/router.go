package presentation

import (
	"JY8752/crawling_app_rest/ent"
	"JY8752/crawling_app_rest/infrastructure/repository"
	"JY8752/crawling_app_rest/presentation/controller"
	"context"
)

func Route(ctx context.Context, client *ent.Client) {
	//crawling route
	cr := repository.NewCrawledUrl(client)
	controller.NewCrawlingController(cr).Register(ctx)
}
