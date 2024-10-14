package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SysMenu holds the schema definition for the SysMenu entity.
type SysMenu struct {
	ent.Schema
}

func (SysMenu) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		ResMixin{},
	}
}

// Fields of the SysMenu.
func (SysMenu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().MaxLen(255).NotEmpty(),
		field.Int64("parent_id").Default(0),
		field.Bool("hidden").Default(false),
	}
}

// Edges of the SysMenu.
func (SysMenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("create_by", SysUser.Type).
			Field("creator").Unique().Required(),
		edge.To("update_by", SysUser.Type).
			Field("updater").Unique().Required(),
		edge.To("delete_by", SysUser.Type).
			Field("deleter").Unique().Required(),
	}
}
