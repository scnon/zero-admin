package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin holds the schema definition for the BaseMixin entity.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Positive(),
		field.Int64("created_at").Default(time.Now().Unix()).Immutable(),
		field.Int64("updated_at").UpdateDefault(time.Now().Unix()),
		field.Int64("deleted_at").Optional(),
		field.Bool("is_deleted").Default(false),
	}
}
