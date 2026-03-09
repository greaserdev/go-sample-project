package schema

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("id"),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{
				"postgres": "timestamptz",
			}).
			Annotations(&entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				"postgres": "timestamptz",
			}).
			Annotations(&entsql.Annotation{Default: "CURRENT_TIMESTAMP"}),
		field.Time("deleted_at").
			Optional().
			Nillable().
			SchemaType(map[string]string{
				"postgres": "timestamptz",
			}),
	}
}

func (b BaseMixin) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.TraverseFunc(func(ctx context.Context, q ent.Query) error {
			// Memanggil fungsi P secara langsung
			if w, ok := q.(interface{ WhereP(...func(*sql.Selector)) }); ok {
				b.P(w)
			}
			return nil
		}),
	}
}

func (BaseMixin) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C("deleted_at")))
		},
	)
}
