// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"reader/internal/book/data/ent/book"
	"reader/internal/book/data/ent/predicate"
	"sync"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeBook = "Book"
)

// BookMutation represents an operation that mutate the Books
// nodes in the graph.
type BookMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	name          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Book, error)
	predicates    []predicate.Book
}

var _ ent.Mutation = (*BookMutation)(nil)

// bookOption allows to manage the mutation configuration using functional options.
type bookOption func(*BookMutation)

// newBookMutation creates new mutation for Book.
func newBookMutation(c config, op Op, opts ...bookOption) *BookMutation {
	m := &BookMutation{
		config:        c,
		op:            op,
		typ:           TypeBook,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBookID sets the id field of the mutation.
func withBookID(id int64) bookOption {
	return func(m *BookMutation) {
		var (
			err   error
			once  sync.Once
			value *Book
		)
		m.oldValue = func(ctx context.Context) (*Book, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Book.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBook sets the old Book of the mutation.
func withBook(node *Book) bookOption {
	return func(m *BookMutation) {
		m.oldValue = func(context.Context) (*Book, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BookMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BookMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Book creation.
func (m *BookMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *BookMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *BookMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *BookMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the Book.
// If the Book object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *BookMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *BookMutation) ResetName() {
	m.name = nil
}

// Op returns the operation name.
func (m *BookMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Book).
func (m *BookMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *BookMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, book.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *BookMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case book.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *BookMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case book.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Book field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *BookMutation) SetField(name string, value ent.Value) error {
	switch name {
	case book.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Book field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *BookMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *BookMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *BookMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Book numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *BookMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *BookMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *BookMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Book nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *BookMutation) ResetField(name string) error {
	switch name {
	case book.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Book field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *BookMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *BookMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *BookMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *BookMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *BookMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *BookMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *BookMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Book unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *BookMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Book edge %s", name)
}
