package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin holds the schema definition for the BaseMixin entity.
type ResMixin struct {
	mixin.Schema
}

// Fields of the BaseMixin.
func (ResMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int8("sort").Default(0).Positive(),
		field.Int64("creator").Positive(),
		field.Int64("updater").Positive(),
		field.Int64("deleter").Positive(),
	}
}
