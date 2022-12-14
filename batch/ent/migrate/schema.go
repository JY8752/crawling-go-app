// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CrawledUrlsColumns holds the columns for the "crawled_urls" table.
	CrawledUrlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// CrawledUrlsTable holds the schema information for the "crawled_urls" table.
	CrawledUrlsTable = &schema.Table{
		Name:       "crawled_urls",
		Columns:    CrawledUrlsColumns,
		PrimaryKey: []*schema.Column{CrawledUrlsColumns[0]},
	}
	// LinkUrlsColumns holds the columns for the "link_urls" table.
	LinkUrlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "referer", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
		{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{"mysql": "datetime"}},
	}
	// LinkUrlsTable holds the schema information for the "link_urls" table.
	LinkUrlsTable = &schema.Table{
		Name:       "link_urls",
		Columns:    LinkUrlsColumns,
		PrimaryKey: []*schema.Column{LinkUrlsColumns[0]},
	}
	// CrawledURLLinkUrlsColumns holds the columns for the "crawled_url_link_urls" table.
	CrawledURLLinkUrlsColumns = []*schema.Column{
		{Name: "crawled_url_id", Type: field.TypeInt},
		{Name: "link_url_id", Type: field.TypeInt},
	}
	// CrawledURLLinkUrlsTable holds the schema information for the "crawled_url_link_urls" table.
	CrawledURLLinkUrlsTable = &schema.Table{
		Name:       "crawled_url_link_urls",
		Columns:    CrawledURLLinkUrlsColumns,
		PrimaryKey: []*schema.Column{CrawledURLLinkUrlsColumns[0], CrawledURLLinkUrlsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "crawled_url_link_urls_crawled_url_id",
				Columns:    []*schema.Column{CrawledURLLinkUrlsColumns[0]},
				RefColumns: []*schema.Column{CrawledUrlsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "crawled_url_link_urls_link_url_id",
				Columns:    []*schema.Column{CrawledURLLinkUrlsColumns[1]},
				RefColumns: []*schema.Column{LinkUrlsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CrawledUrlsTable,
		LinkUrlsTable,
		CrawledURLLinkUrlsTable,
	}
)

func init() {
	CrawledURLLinkUrlsTable.ForeignKeys[0].RefTable = CrawledUrlsTable
	CrawledURLLinkUrlsTable.ForeignKeys[1].RefTable = LinkUrlsTable
}
