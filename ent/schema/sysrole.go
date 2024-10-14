package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SysRole holds the schema definition for the SysRole entity.
type SysRole struct {
	ent.Schema
}

func (SysRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ResMixin{},
	}
}

// Fields of the SysRole.
func (SysRole) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(255).NotEmpty().Unique(),
		field.Int8("status").Default(1),
		field.String("remark").MaxLen(512),
	}
}

// Edges of the SysRole.
func (SysRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("create_by", SysUser.Type).
			Field("creator").Unique().Required(),
		edge.To("update_by", SysUser.Type).
			Field("updater").Unique().Required(),
		edge.To("delete_by", SysUser.Type).
			Field("deleter").Unique().Required(),
	}
}
