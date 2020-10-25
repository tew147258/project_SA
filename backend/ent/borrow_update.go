// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/tew147258/app/ent/borrow"
	"github.com/tew147258/app/ent/confirmation"
	"github.com/tew147258/app/ent/predicate"
)

// BorrowUpdate is the builder for updating Borrow entities.
type BorrowUpdate struct {
	config
	hooks      []Hook
	mutation   *BorrowMutation
	predicates []predicate.Borrow
}

// Where adds a new predicate for the builder.
func (bu *BorrowUpdate) Where(ps ...predicate.Borrow) *BorrowUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetType sets the type field.
func (bu *BorrowUpdate) SetType(s string) *BorrowUpdate {
	bu.mutation.SetType(s)
	return bu
}

// AddBorrowConfirmationIDs adds the BorrowConfirmation edge to Confirmation by ids.
func (bu *BorrowUpdate) AddBorrowConfirmationIDs(ids ...int) *BorrowUpdate {
	bu.mutation.AddBorrowConfirmationIDs(ids...)
	return bu
}

// AddBorrowConfirmation adds the BorrowConfirmation edges to Confirmation.
func (bu *BorrowUpdate) AddBorrowConfirmation(c ...*Confirmation) *BorrowUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.AddBorrowConfirmationIDs(ids...)
}

// Mutation returns the BorrowMutation object of the builder.
func (bu *BorrowUpdate) Mutation() *BorrowMutation {
	return bu.mutation
}

// RemoveBorrowConfirmationIDs removes the BorrowConfirmation edge to Confirmation by ids.
func (bu *BorrowUpdate) RemoveBorrowConfirmationIDs(ids ...int) *BorrowUpdate {
	bu.mutation.RemoveBorrowConfirmationIDs(ids...)
	return bu
}

// RemoveBorrowConfirmation removes BorrowConfirmation edges to Confirmation.
func (bu *BorrowUpdate) RemoveBorrowConfirmation(c ...*Confirmation) *BorrowUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.RemoveBorrowConfirmationIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BorrowUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := bu.mutation.GetType(); ok {
		if err := borrow.TypeValidator(v); err != nil {
			return 0, &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BorrowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BorrowUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BorrowUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BorrowUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BorrowUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   borrow.Table,
			Columns: borrow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: borrow.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: borrow.FieldType,
		})
	}
	if nodes := bu.mutation.RemovedBorrowConfirmationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   borrow.BorrowConfirmationTable,
			Columns: []string{borrow.BorrowConfirmationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: confirmation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.BorrowConfirmationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   borrow.BorrowConfirmationTable,
			Columns: []string{borrow.BorrowConfirmationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: confirmation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{borrow.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BorrowUpdateOne is the builder for updating a single Borrow entity.
type BorrowUpdateOne struct {
	config
	hooks    []Hook
	mutation *BorrowMutation
}

// SetType sets the type field.
func (buo *BorrowUpdateOne) SetType(s string) *BorrowUpdateOne {
	buo.mutation.SetType(s)
	return buo
}

// AddBorrowConfirmationIDs adds the BorrowConfirmation edge to Confirmation by ids.
func (buo *BorrowUpdateOne) AddBorrowConfirmationIDs(ids ...int) *BorrowUpdateOne {
	buo.mutation.AddBorrowConfirmationIDs(ids...)
	return buo
}

// AddBorrowConfirmation adds the BorrowConfirmation edges to Confirmation.
func (buo *BorrowUpdateOne) AddBorrowConfirmation(c ...*Confirmation) *BorrowUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.AddBorrowConfirmationIDs(ids...)
}

// Mutation returns the BorrowMutation object of the builder.
func (buo *BorrowUpdateOne) Mutation() *BorrowMutation {
	return buo.mutation
}

// RemoveBorrowConfirmationIDs removes the BorrowConfirmation edge to Confirmation by ids.
func (buo *BorrowUpdateOne) RemoveBorrowConfirmationIDs(ids ...int) *BorrowUpdateOne {
	buo.mutation.RemoveBorrowConfirmationIDs(ids...)
	return buo
}

// RemoveBorrowConfirmation removes BorrowConfirmation edges to Confirmation.
func (buo *BorrowUpdateOne) RemoveBorrowConfirmation(c ...*Confirmation) *BorrowUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.RemoveBorrowConfirmationIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (buo *BorrowUpdateOne) Save(ctx context.Context) (*Borrow, error) {
	if v, ok := buo.mutation.GetType(); ok {
		if err := borrow.TypeValidator(v); err != nil {
			return nil, &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}

	var (
		err  error
		node *Borrow
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BorrowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BorrowUpdateOne) SaveX(ctx context.Context) *Borrow {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BorrowUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BorrowUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BorrowUpdateOne) sqlSave(ctx context.Context) (b *Borrow, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   borrow.Table,
			Columns: borrow.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: borrow.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Borrow.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: borrow.FieldType,
		})
	}
	if nodes := buo.mutation.RemovedBorrowConfirmationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   borrow.BorrowConfirmationTable,
			Columns: []string{borrow.BorrowConfirmationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: confirmation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.BorrowConfirmationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   borrow.BorrowConfirmationTable,
			Columns: []string{borrow.BorrowConfirmationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: confirmation.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	b = &Borrow{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{borrow.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}
