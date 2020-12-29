// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Piichet/app/ent/title"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Title is the model entity for the Title schema.
type Title struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Title) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // title
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Title fields.
func (t *Title) assignValues(values ...interface{}) error {
	if m, n := len(values), len(title.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[0])
	} else if value.Valid {
		t.Title = value.String
	}
	return nil
}

// Update returns a builder for updating this Title.
// Note that, you need to call Title.Unwrap() before calling this method, if this Title
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Title) Update() *TitleUpdateOne {
	return (&TitleClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Title) Unwrap() *Title {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Title is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Title) String() string {
	var builder strings.Builder
	builder.WriteString("Title(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", title=")
	builder.WriteString(t.Title)
	builder.WriteByte(')')
	return builder.String()
}

// Titles is a parsable slice of Title.
type Titles []*Title

func (t Titles) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}