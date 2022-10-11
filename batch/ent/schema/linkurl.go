package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// LinkUrl holds the schema definition for the LinkUrl entity.
type LinkUrl struct {
	ent.Schema
}

// Fields of the LinkUrl.
func (LinkUrl) Fields() []ent.Field {
	return []ent.Field{
		field.String("url").NotEmpty().SchemaType(map[string]string{
			dialect.MySQL: "text",
		}),
		field.String("referer").NotEmpty().SchemaType(map[string]string{
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

// Edges of the LinkUrl.
func (LinkUrl) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("base_url", CrawledUrl.Type).
			Ref("link_urls"),
	}
}
