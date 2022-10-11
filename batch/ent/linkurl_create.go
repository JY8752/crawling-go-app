// Code generated by ent, DO NOT EDIT.

package ent

import (
	"JY8752/crawling_app_batch/ent/crawledurl"
	"JY8752/crawling_app_batch/ent/linkurl"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkUrlCreate is the builder for creating a LinkUrl entity.
type LinkUrlCreate struct {
	config
	mutation *LinkUrlMutation
	hooks    []Hook
}

// SetURL sets the "url" field.
func (luc *LinkUrlCreate) SetURL(s string) *LinkUrlCreate {
	luc.mutation.SetURL(s)
	return luc
}

// SetReferer sets the "referer" field.
func (luc *LinkUrlCreate) SetReferer(s string) *LinkUrlCreate {
	luc.mutation.SetReferer(s)
	return luc
}

// SetCreatedAt sets the "created_at" field.
func (luc *LinkUrlCreate) SetCreatedAt(t time.Time) *LinkUrlCreate {
	luc.mutation.SetCreatedAt(t)
	return luc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luc *LinkUrlCreate) SetNillableCreatedAt(t *time.Time) *LinkUrlCreate {
	if t != nil {
		luc.SetCreatedAt(*t)
	}
	return luc
}

// SetUpdatedAt sets the "updated_at" field.
func (luc *LinkUrlCreate) SetUpdatedAt(t time.Time) *LinkUrlCreate {
	luc.mutation.SetUpdatedAt(t)
	return luc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (luc *LinkUrlCreate) SetNillableUpdatedAt(t *time.Time) *LinkUrlCreate {
	if t != nil {
		luc.SetUpdatedAt(*t)
	}
	return luc
}

// AddBaseURLIDs adds the "base_url" edge to the CrawledUrl entity by IDs.
func (luc *LinkUrlCreate) AddBaseURLIDs(ids ...int) *LinkUrlCreate {
	luc.mutation.AddBaseURLIDs(ids...)
	return luc
}

// AddBaseURL adds the "base_url" edges to the CrawledUrl entity.
func (luc *LinkUrlCreate) AddBaseURL(c ...*CrawledUrl) *LinkUrlCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return luc.AddBaseURLIDs(ids...)
}

// Mutation returns the LinkUrlMutation object of the builder.
func (luc *LinkUrlCreate) Mutation() *LinkUrlMutation {
	return luc.mutation
}

// Save creates the LinkUrl in the database.
func (luc *LinkUrlCreate) Save(ctx context.Context) (*LinkUrl, error) {
	var (
		err  error
		node *LinkUrl
	)
	luc.defaults()
	if len(luc.hooks) == 0 {
		if err = luc.check(); err != nil {
			return nil, err
		}
		node, err = luc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinkUrlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = luc.check(); err != nil {
				return nil, err
			}
			luc.mutation = mutation
			if node, err = luc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(luc.hooks) - 1; i >= 0; i-- {
			if luc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, luc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*LinkUrl)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LinkUrlMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (luc *LinkUrlCreate) SaveX(ctx context.Context) *LinkUrl {
	v, err := luc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (luc *LinkUrlCreate) Exec(ctx context.Context) error {
	_, err := luc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luc *LinkUrlCreate) ExecX(ctx context.Context) {
	if err := luc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luc *LinkUrlCreate) defaults() {
	if _, ok := luc.mutation.CreatedAt(); !ok {
		v := linkurl.DefaultCreatedAt
		luc.mutation.SetCreatedAt(v)
	}
	if _, ok := luc.mutation.UpdatedAt(); !ok {
		v := linkurl.DefaultUpdatedAt
		luc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (luc *LinkUrlCreate) check() error {
	if _, ok := luc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "LinkUrl.url"`)}
	}
	if v, ok := luc.mutation.URL(); ok {
		if err := linkurl.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "LinkUrl.url": %w`, err)}
		}
	}
	if _, ok := luc.mutation.Referer(); !ok {
		return &ValidationError{Name: "referer", err: errors.New(`ent: missing required field "LinkUrl.referer"`)}
	}
	if v, ok := luc.mutation.Referer(); ok {
		if err := linkurl.RefererValidator(v); err != nil {
			return &ValidationError{Name: "referer", err: fmt.Errorf(`ent: validator failed for field "LinkUrl.referer": %w`, err)}
		}
	}
	if _, ok := luc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "LinkUrl.created_at"`)}
	}
	if _, ok := luc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "LinkUrl.updated_at"`)}
	}
	return nil
}

func (luc *LinkUrlCreate) sqlSave(ctx context.Context) (*LinkUrl, error) {
	_node, _spec := luc.createSpec()
	if err := sqlgraph.CreateNode(ctx, luc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (luc *LinkUrlCreate) createSpec() (*LinkUrl, *sqlgraph.CreateSpec) {
	var (
		_node = &LinkUrl{config: luc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: linkurl.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: linkurl.FieldID,
			},
		}
	)
	if value, ok := luc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linkurl.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := luc.mutation.Referer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: linkurl.FieldReferer,
		})
		_node.Referer = value
	}
	if value, ok := luc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: linkurl.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := luc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: linkurl.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := luc.mutation.BaseURLIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   linkurl.BaseURLTable,
			Columns: linkurl.BaseURLPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: crawledurl.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LinkUrlCreateBulk is the builder for creating many LinkUrl entities in bulk.
type LinkUrlCreateBulk struct {
	config
	builders []*LinkUrlCreate
}

// Save creates the LinkUrl entities in the database.
func (lucb *LinkUrlCreateBulk) Save(ctx context.Context) ([]*LinkUrl, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lucb.builders))
	nodes := make([]*LinkUrl, len(lucb.builders))
	mutators := make([]Mutator, len(lucb.builders))
	for i := range lucb.builders {
		func(i int, root context.Context) {
			builder := lucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LinkUrlMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lucb *LinkUrlCreateBulk) SaveX(ctx context.Context) []*LinkUrl {
	v, err := lucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lucb *LinkUrlCreateBulk) Exec(ctx context.Context) error {
	_, err := lucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lucb *LinkUrlCreateBulk) ExecX(ctx context.Context) {
	if err := lucb.Exec(ctx); err != nil {
		panic(err)
	}
}
