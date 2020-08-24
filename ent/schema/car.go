package schema

import (
	"github.com/facebook/ent"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return nil

}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return nil
}
