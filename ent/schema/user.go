package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {

	return []ent.Field{
		field.Int("age").Positive(),
		field.String("name"),
		field.String("username").Unique(),
		field.String("email").MaxLen(40),
		field.String("password").MaxLen(128),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", Group.Type),
		edge.To("friends", User.Type),
		//edge.From("groups", Group.Type).Ref("users"),
	}
}

func (User) Index() []ent.Index {
	return []ent.Index{
		index.Fields("age", "name").Unique(),
	}
}
