// Code generated by ent, DO NOT EDIT.

package ent

import (
	"JY8752/crawling_app_batch/ent/linkurl"
	"JY8752/crawling_app_batch/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LinkUrlDelete is the builder for deleting a LinkUrl entity.
type LinkUrlDelete struct {
	config
	hooks    []Hook
	mutation *LinkUrlMutation
}

// Where appends a list predicates to the LinkUrlDelete builder.
func (lud *LinkUrlDelete) Where(ps ...predicate.LinkUrl) *LinkUrlDelete {
	lud.mutation.Where(ps...)
	return lud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (lud *LinkUrlDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lud.hooks) == 0 {
		affected, err = lud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LinkUrlMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lud.mutation = mutation
			affected, err = lud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lud.hooks) - 1; i >= 0; i-- {
			if lud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (lud *LinkUrlDelete) ExecX(ctx context.Context) int {
	n, err := lud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (lud *LinkUrlDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: linkurl.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: linkurl.FieldID,
			},
		},
	}
	if ps := lud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, lud.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// LinkUrlDeleteOne is the builder for deleting a single LinkUrl entity.
type LinkUrlDeleteOne struct {
	lud *LinkUrlDelete
}

// Exec executes the deletion query.
func (ludo *LinkUrlDeleteOne) Exec(ctx context.Context) error {
	n, err := ludo.lud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{linkurl.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ludo *LinkUrlDeleteOne) ExecX(ctx context.Context) {
	ludo.lud.ExecX(ctx)
}
