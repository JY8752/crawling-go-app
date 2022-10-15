package data

import (
	"context"
	"testing"
	"time"

	"JY8752/crawling_app_batch/ent"
	"JY8752/crawling_app_batch/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
)

func TestSave(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	cu := NewCrawledUrl(client)

	ctx := context.Background()
	// tt := time.Now()
	tt := time.Date(2022, 10, 15, 0, 0, 0, 0, time.Local)
	saved, err := cu.Save(ctx, "https://test.com", &tt)

	if err != nil {
		t.Errorf("failed save url. err: %v\n", err.Error())
	}

	find := cu.FindById(ctx, saved.ID)

	assertCrawledUrl(t, saved, find)
}

func assertCrawledUrl(t *testing.T, expect, act *ent.CrawledUrl) {
	t.Helper()
	if expect.ID != act.ID {
		t.Errorf("assert error id. expected: %v, act: %v\n", expect.ID, act.ID)
	}
	if expect.URL != act.URL {
		t.Errorf("assert error url. expected: %v, act: %v\n", expect.URL, act.URL)
	}
	if !expect.UpdatedAt.Equal(act.UpdatedAt) {
		t.Errorf("assert error updated_at. expected: %v, act: %v\n", expect.UpdatedAt, act.UpdatedAt)
	}
	if !expect.CreatedAt.Equal(act.CreatedAt) {
		t.Errorf("assert error created_at.  expected: %v, act: %v\n", expect.CreatedAt, act.CreatedAt)
	}
}
