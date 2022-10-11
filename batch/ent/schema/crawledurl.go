package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CrawledUrl holds the schema definition for the CrawledUrl entity.
type CrawledUrl struct {
	ent.Schema
}

// Fields of the CrawledUrl.
func (CrawledUrl) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").NotEmpty().SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.Time("created_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp",
		}).Default(time.Now()),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp",
		}).Default(time.Now()),
	}
}

// Edges of the CrawledUrl.
func (CrawledUrl) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("link_urls", LinkUrl.Type),
	}
}
