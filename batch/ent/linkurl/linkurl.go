// Code generated by ent, DO NOT EDIT.

package linkurl

import (
	"time"
)

const (
	// Label holds the string label denoting the linkurl type in the database.
	Label = "link_url"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldReferer holds the string denoting the referer field in the database.
	FieldReferer = "referer"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeBaseURL holds the string denoting the base_url edge name in mutations.
	EdgeBaseURL = "base_url"
	// Table holds the table name of the linkurl in the database.
	Table = "link_urls"
	// BaseURLTable is the table that holds the base_url relation/edge. The primary key declared below.
	BaseURLTable = "crawled_url_link_urls"
	// BaseURLInverseTable is the table name for the CrawledUrl entity.
	// It exists in this package in order to avoid circular dependency with the "crawledurl" package.
	BaseURLInverseTable = "crawled_urls"
)

// Columns holds all SQL columns for linkurl fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldReferer,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// BaseURLPrimaryKey and BaseURLColumn2 are the table columns denoting the
	// primary key for the base_url relation (M2M).
	BaseURLPrimaryKey = []string{"crawled_url_id", "link_url_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// RefererValidator is a validator for the "referer" field. It is called by the builders before save.
	RefererValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt time.Time
)
