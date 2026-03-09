package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Credentials holds the schema definition for the Credentials entity.
type Credentials struct {
	ent.Schema
}

// Fields of the Credentials.
func (Credentials) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").NotEmpty(),
		field.String("password_hash").Sensitive().NotEmpty(),
		field.Time("email_verified_at").SchemaType(map[string]string{
			"postgres": "timestamptz",
		}).Nillable().Optional(),
	}
}

// Edges of the Credentials.
func (Credentials) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique(),
	}
}

func (Credentials) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
