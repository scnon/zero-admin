package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SysUser holds the schema definition for the SysUser entity.
type SysUser struct {
	ent.Schema
}

func (SysUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ResMixin{},
	}
}

// Fields of the SysUser.
func (SysUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique().NotEmpty().MaxLen(255),
		field.String("password").NotEmpty().MaxLen(255),
		field.String("nickname").NotEmpty().MaxLen(255),
		field.String("avatar").MaxLen(255),
		field.Int8("status").Default(1),
		field.String("remark").MaxLen(512),
	}
}

// Edges of the SysUser.
func (SysUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("create_by", SysUser.Type).
			Field("creator").Unique().Required(),
		edge.To("update_by", SysUser.Type).
			Field("updater").Unique().Required(),
		edge.To("delete_by", SysUser.Type).
			Field("deleter").Unique().Required(),
	}
}
