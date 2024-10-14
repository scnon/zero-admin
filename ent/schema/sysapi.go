package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SysApi holds the schema definition for the SysApi entity.
type SysApi struct {
	ent.Schema
}

func (SysApi) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ResMixin{},
	}
}

// Fields of the SysApi.
func (SysApi) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().MaxLen(255).NotEmpty(),
		field.String("path").Unique().MaxLen(255).NotEmpty(),
	}
}

// Edges of the SysApi.
func (SysApi) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("create_by", SysUser.Type).
			Field("creator").Unique().Required(),
		edge.To("update_by", SysUser.Type).
			Field("updater").Unique().Required(),
		edge.To("delete_by", SysUser.Type).
			Field("deleter").Unique().Required(),
	}
}
