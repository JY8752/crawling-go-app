package data

import (
	data "JY8752/crawling_app_batch/data/crawledurl"
	"JY8752/crawling_app_batch/ent"
	"JY8752/crawling_app_batch/ent/enttest"
	"context"
	"fmt"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestLinkUrl(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	lu := NewLinkUrl(client)

	ctx := context.Background()
	tt := time.Date(2022, 10, 15, 0, 0, 0, 0, time.Local)

	cu, err := data.NewCrawledUrl(client).Save(ctx, "https://test.com", &tt)
	if err != nil {
		t.Errorf("failed save crawled url. err: %v\n", err.Error())
	}

	l1 := lu.Save(ctx, "https://test.com/page1", "https://test,com", &tt, cu)
	l2 := lu.Save(ctx, "https://test.com/page2", "https://test,com", &tt, cu)

	find := lu.FindByCrawledUrlId(ctx, cu.ID)

	fmt.Printf("link1: %v link2: %v find: %v\n", l1, l2, find)

	if len(find) != 2 {
		t.Errorf("expect link_url 2 record, but %d\n", len(find))
	}

	assertLinkUrl(t, l1, find[0])
	assertLinkUrl(t, l2, find[1])
}

func assertLinkUrl(t *testing.T, expect, act *ent.LinkUrl) {
	t.Helper()

	if expect.ID != act.ID {
		t.Errorf("assert error id. expected: %v, act: %v\n", expect.ID, act.ID)
		return
	}

	if expect.URL != act.URL {
		t.Errorf("assert error url. expected: %v, act: %v\n", expect.URL, act.URL)
		return
	}

	if expect.Referer != act.Referer {
		t.Errorf("assert error referer. expected: %v, act: %v\n", expect.Referer, act.Referer)
		return
	}

	if !expect.UpdatedAt.Equal(act.UpdatedAt) {
		t.Errorf("assert error updated_at. expected: %v, act: %v\n", expect.UpdatedAt, act.UpdatedAt)
		return
	}

	if !expect.CreatedAt.Equal(act.CreatedAt) {
		t.Errorf("assert error created_at. expected: %v, act: %v\n", expect.CreatedAt, act.CreatedAt)
		return
	}
}
