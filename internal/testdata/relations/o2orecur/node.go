// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Node struct {
	ent.Schema
}

func (Node) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.Int("value"), field.Int("node_next").Optional()}
}
func (Node) Edges() []ent.Edge {
	return []ent.Edge{edge.To("child_node", Node.Type).Unique(), edge.From("parent_node", Node.Type).Unique().Field("node_next")}
}
func (Node) Annotations() []schema.Annotation {
	return nil
}