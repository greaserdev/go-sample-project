package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().NotEmpty(),
		field.String("first_name").NotEmpty().MaxLen(100),
		field.String("last_name").NotEmpty().MaxLen(100),
		field.Time("verified_at").SchemaType(map[string]string{
			"postgres": "timestamptz",
		}).Optional().Nillable(),
		field.String("avatar_url").Nillable().Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("credential", Credentials.Type).Ref("user").Unique(),
	}
}

// Mixin of the User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
