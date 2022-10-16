// Code generated by ent, DO NOT EDIT.

package ent

import (
	"JY8752/crawling_app_rest/ent/crawledurl"
	"JY8752/crawling_app_rest/ent/linkurl"
	"JY8752/crawling_app_rest/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkUrlQuery is the builder for querying LinkUrl entities.
type LinkUrlQuery struct {
	config
	limit       *int
	offset      *int
	unique      *bool
	order       []OrderFunc
	fields      []string
	predicates  []predicate.LinkUrl
	withBaseURL *CrawledUrlQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LinkUrlQuery builder.
func (luq *LinkUrlQuery) Where(ps ...predicate.LinkUrl) *LinkUrlQuery {
	luq.predicates = append(luq.predicates, ps...)
	return luq
}

// Limit adds a limit step to the query.
func (luq *LinkUrlQuery) Limit(limit int) *LinkUrlQuery {
	luq.limit = &limit
	return luq
}

// Offset adds an offset step to the query.
func (luq *LinkUrlQuery) Offset(offset int) *LinkUrlQuery {
	luq.offset = &offset
	return luq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (luq *LinkUrlQuery) Unique(unique bool) *LinkUrlQuery {
	luq.unique = &unique
	return luq
}

// Order adds an order step to the query.
func (luq *LinkUrlQuery) Order(o ...OrderFunc) *LinkUrlQuery {
	luq.order = append(luq.order, o...)
	return luq
}

// QueryBaseURL chains the current query on the "base_url" edge.
func (luq *LinkUrlQuery) QueryBaseURL() *CrawledUrlQuery {
	query := &CrawledUrlQuery{config: luq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := luq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := luq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(linkurl.Table, linkurl.FieldID, selector),
			sqlgraph.To(crawledurl.Table, crawledurl.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, linkurl.BaseURLTable, linkurl.BaseURLPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(luq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first LinkUrl entity from the query.
// Returns a *NotFoundError when no LinkUrl was found.
func (luq *LinkUrlQuery) First(ctx context.Context) (*LinkUrl, error) {
	nodes, err := luq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{linkurl.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (luq *LinkUrlQuery) FirstX(ctx context.Context) *LinkUrl {
	node, err := luq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LinkUrl ID from the query.
// Returns a *NotFoundError when no LinkUrl ID was found.
func (luq *LinkUrlQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = luq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{linkurl.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (luq *LinkUrlQuery) FirstIDX(ctx context.Context) int {
	id, err := luq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LinkUrl entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LinkUrl entity is found.
// Returns a *NotFoundError when no LinkUrl entities are found.
func (luq *LinkUrlQuery) Only(ctx context.Context) (*LinkUrl, error) {
	nodes, err := luq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{linkurl.Label}
	default:
		return nil, &NotSingularError{linkurl.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (luq *LinkUrlQuery) OnlyX(ctx context.Context) *LinkUrl {
	node, err := luq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LinkUrl ID in the query.
// Returns a *NotSingularError when more than one LinkUrl ID is found.
// Returns a *NotFoundError when no entities are found.
func (luq *LinkUrlQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = luq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{linkurl.Label}
	default:
		err = &NotSingularError{linkurl.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (luq *LinkUrlQuery) OnlyIDX(ctx context.Context) int {
	id, err := luq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LinkUrls.
func (luq *LinkUrlQuery) All(ctx context.Context) ([]*LinkUrl, error) {
	if err := luq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return luq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (luq *LinkUrlQuery) AllX(ctx context.Context) []*LinkUrl {
	nodes, err := luq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LinkUrl IDs.
func (luq *LinkUrlQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := luq.Select(linkurl.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (luq *LinkUrlQuery) IDsX(ctx context.Context) []int {
	ids, err := luq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (luq *LinkUrlQuery) Count(ctx context.Context) (int, error) {
	if err := luq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return luq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (luq *LinkUrlQuery) CountX(ctx context.Context) int {
	count, err := luq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (luq *LinkUrlQuery) Exist(ctx context.Context) (bool, error) {
	if err := luq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return luq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (luq *LinkUrlQuery) ExistX(ctx context.Context) bool {
	exist, err := luq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LinkUrlQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (luq *LinkUrlQuery) Clone() *LinkUrlQuery {
	if luq == nil {
		return nil
	}
	return &LinkUrlQuery{
		config:      luq.config,
		limit:       luq.limit,
		offset:      luq.offset,
		order:       append([]OrderFunc{}, luq.order...),
		predicates:  append([]predicate.LinkUrl{}, luq.predicates...),
		withBaseURL: luq.withBaseURL.Clone(),
		// clone intermediate query.
		sql:    luq.sql.Clone(),
		path:   luq.path,
		unique: luq.unique,
	}
}

// WithBaseURL tells the query-builder to eager-load the nodes that are connected to
// the "base_url" edge. The optional arguments are used to configure the query builder of the edge.
func (luq *LinkUrlQuery) WithBaseURL(opts ...func(*CrawledUrlQuery)) *LinkUrlQuery {
	query := &CrawledUrlQuery{config: luq.config}
	for _, opt := range opts {
		opt(query)
	}
	luq.withBaseURL = query
	return luq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		URL string `json:"url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LinkUrl.Query().
//		GroupBy(linkurl.FieldURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (luq *LinkUrlQuery) GroupBy(field string, fields ...string) *LinkUrlGroupBy {
	grbuild := &LinkUrlGroupBy{config: luq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := luq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return luq.sqlQuery(ctx), nil
	}
	grbuild.label = linkurl.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		URL string `json:"url,omitempty"`
//	}
//
//	client.LinkUrl.Query().
//		Select(linkurl.FieldURL).
//		Scan(ctx, &v)
//
func (luq *LinkUrlQuery) Select(fields ...string) *LinkUrlSelect {
	luq.fields = append(luq.fields, fields...)
	selbuild := &LinkUrlSelect{LinkUrlQuery: luq}
	selbuild.label = linkurl.Label
	selbuild.flds, selbuild.scan = &luq.fields, selbuild.Scan
	return selbuild
}

func (luq *LinkUrlQuery) prepareQuery(ctx context.Context) error {
	for _, f := range luq.fields {
		if !linkurl.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if luq.path != nil {
		prev, err := luq.path(ctx)
		if err != nil {
			return err
		}
		luq.sql = prev
	}
	return nil
}

func (luq *LinkUrlQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LinkUrl, error) {
	var (
		nodes       = []*LinkUrl{}
		_spec       = luq.querySpec()
		loadedTypes = [1]bool{
			luq.withBaseURL != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*LinkUrl).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &LinkUrl{config: luq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, luq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := luq.withBaseURL; query != nil {
		if err := luq.loadBaseURL(ctx, query, nodes,
			func(n *LinkUrl) { n.Edges.BaseURL = []*CrawledUrl{} },
			func(n *LinkUrl, e *CrawledUrl) { n.Edges.BaseURL = append(n.Edges.BaseURL, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (luq *LinkUrlQuery) loadBaseURL(ctx context.Context, query *CrawledUrlQuery, nodes []*LinkUrl, init func(*LinkUrl), assign func(*LinkUrl, *CrawledUrl)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*LinkUrl)
	nids := make(map[int]map[*LinkUrl]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(linkurl.BaseURLTable)
		s.Join(joinT).On(s.C(crawledurl.FieldID), joinT.C(linkurl.BaseURLPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(linkurl.BaseURLPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(linkurl.BaseURLPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*LinkUrl]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "base_url" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (luq *LinkUrlQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := luq.querySpec()
	_spec.Node.Columns = luq.fields
	if len(luq.fields) > 0 {
		_spec.Unique = luq.unique != nil && *luq.unique
	}
	return sqlgraph.CountNodes(ctx, luq.driver, _spec)
}

func (luq *LinkUrlQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := luq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (luq *LinkUrlQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   linkurl.Table,
			Columns: linkurl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: linkurl.FieldID,
			},
		},
		From:   luq.sql,
		Unique: true,
	}
	if unique := luq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := luq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, linkurl.FieldID)
		for i := range fields {
			if fields[i] != linkurl.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := luq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := luq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := luq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := luq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (luq *LinkUrlQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(luq.driver.Dialect())
	t1 := builder.Table(linkurl.Table)
	columns := luq.fields
	if len(columns) == 0 {
		columns = linkurl.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if luq.sql != nil {
		selector = luq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if luq.unique != nil && *luq.unique {
		selector.Distinct()
	}
	for _, p := range luq.predicates {
		p(selector)
	}
	for _, p := range luq.order {
		p(selector)
	}
	if offset := luq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := luq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LinkUrlGroupBy is the group-by builder for LinkUrl entities.
type LinkUrlGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lugb *LinkUrlGroupBy) Aggregate(fns ...AggregateFunc) *LinkUrlGroupBy {
	lugb.fns = append(lugb.fns, fns...)
	return lugb
}

// Scan applies the group-by query and scans the result into the given value.
func (lugb *LinkUrlGroupBy) Scan(ctx context.Context, v any) error {
	query, err := lugb.path(ctx)
	if err != nil {
		return err
	}
	lugb.sql = query
	return lugb.sqlScan(ctx, v)
}

func (lugb *LinkUrlGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range lugb.fields {
		if !linkurl.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lugb *LinkUrlGroupBy) sqlQuery() *sql.Selector {
	selector := lugb.sql.Select()
	aggregation := make([]string, 0, len(lugb.fns))
	for _, fn := range lugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lugb.fields)+len(lugb.fns))
		for _, f := range lugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lugb.fields...)...)
}

// LinkUrlSelect is the builder for selecting fields of LinkUrl entities.
type LinkUrlSelect struct {
	*LinkUrlQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (lus *LinkUrlSelect) Scan(ctx context.Context, v any) error {
	if err := lus.prepareQuery(ctx); err != nil {
		return err
	}
	lus.sql = lus.LinkUrlQuery.sqlQuery(ctx)
	return lus.sqlScan(ctx, v)
}

func (lus *LinkUrlSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := lus.sql.Query()
	if err := lus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}