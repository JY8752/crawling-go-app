package presentation

import (
	"JY8752/crawling_app_rest/ent"
	"JY8752/crawling_app_rest/infrastructure/repository"
	"JY8752/crawling_app_rest/presentation/controller"
	"context"

	"github.com/gorilla/mux"
)

func Route(ctx context.Context, client *ent.Client, r *mux.Router) {
	//crawling route
	cr := repository.NewCrawledUrl(client)
	lu := repository.NewLinkUrl(client)
	controller.NewCrawlingController(cr, lu).Register(ctx, r)
}
