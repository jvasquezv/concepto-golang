// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo/todoadapter/ent/predicate"
	"github.com/sagikazarmark/modern-go-application/internal/app/mga/todo/todoadapter/ent/todoitem"
)

// TodoItemDelete is the builder for deleting a TodoItem entity.
type TodoItemDelete struct {
	config
	hooks      []Hook
	mutation   *TodoItemMutation
	predicates []predicate.TodoItem
}

// Where adds a new predicate to the delete builder.
func (tid *TodoItemDelete) Where(ps ...predicate.TodoItem) *TodoItemDelete {
	tid.predicates = append(tid.predicates, ps...)
	return tid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tid *TodoItemDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tid.hooks) == 0 {
		affected, err = tid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TodoItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tid.mutation = mutation
			affected, err = tid.sqlExec(ctx)
			return affected, err
		})
		for i := len(tid.hooks) - 1; i >= 0; i-- {
			mut = tid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tid *TodoItemDelete) ExecX(ctx context.Context) int {
	n, err := tid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tid *TodoItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: todoitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todoitem.FieldID,
			},
		},
	}
	if ps := tid.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tid.driver, _spec)
}

// TodoItemDeleteOne is the builder for deleting a single TodoItem entity.
type TodoItemDeleteOne struct {
	tid *TodoItemDelete
}

// Exec executes the deletion query.
func (tido *TodoItemDeleteOne) Exec(ctx context.Context) error {
	n, err := tido.tid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{todoitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tido *TodoItemDeleteOne) ExecX(ctx context.Context) {
	tido.tid.ExecX(ctx)
}
