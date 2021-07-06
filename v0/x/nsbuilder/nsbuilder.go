// Package nsbuilder implements a builder-pattern for defining Authzed
// Namespaces.
package nsbuilder

import v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"

// Namespace creates a namespace definition with one or more defined relations.
func Namespace(name string, relations ...*v0.Relation) *v0.NamespaceDefinition {
	return &v0.NamespaceDefinition{
		Name:     name,
		Relation: relations,
	}
}

// Relation creates a relation definition with an optional rewrite definition.
func Relation(name string, rewrite *v0.UsersetRewrite) *v0.Relation {
	return &v0.Relation{
		Name:           name,
		UsersetRewrite: rewrite,
	}
}

// Union creates a rewrite definition that combines/considers usersets in all children.
func Union(firstChild *v0.SetOperation_Child, rest ...*v0.SetOperation_Child) *v0.UsersetRewrite {
	return &v0.UsersetRewrite{
		RewriteOperation: &v0.UsersetRewrite_Union{
			Union: setOperation(firstChild, rest),
		},
	}
}

// Intersection creates a rewrite definition that returns/considers only usersets present in all children.
func Intersection(firstChild *v0.SetOperation_Child, rest ...*v0.SetOperation_Child) *v0.UsersetRewrite {
	return &v0.UsersetRewrite{
		RewriteOperation: &v0.UsersetRewrite_Intersection{
			Intersection: setOperation(firstChild, rest),
		},
	}
}

// Exclusion creates a rewrite definition that starts with the usersets of the first child
// and iteratively removes usersets that appear in the remaining children.
func Exclusion(firstChild *v0.SetOperation_Child, rest ...*v0.SetOperation_Child) *v0.UsersetRewrite {
	return &v0.UsersetRewrite{
		RewriteOperation: &v0.UsersetRewrite_Exclusion{
			Exclusion: setOperation(firstChild, rest),
		},
	}
}

func setOperation(firstChild *v0.SetOperation_Child, rest []*v0.SetOperation_Child) *v0.SetOperation {
	children := append([]*v0.SetOperation_Child{firstChild}, rest...)
	return &v0.SetOperation{
		Child: children,
	}
}

// This creates a child for a set operation that references only direct usersets with the parent relation.
func This() *v0.SetOperation_Child {
	return &v0.SetOperation_Child{
		ChildType: &v0.SetOperation_Child_XThis{},
	}
}

// ComputesUserset creates a child for a set operation that follows a relation on the given starting object.
func ComputedUserset(relation string) *v0.SetOperation_Child {
	return &v0.SetOperation_Child{
		ChildType: &v0.SetOperation_Child_ComputedUserset{
			ComputedUserset: &v0.ComputedUserset{
				Relation: relation,
			},
		},
	}
}

// TupleToUserset creates a child which first loads all tuples with the specific relation,
// and then unions all children on the usersets found by following a relation on those loaded
// tuples.
func TupleToUserset(tuplesetRelation, usersetRelation string) *v0.SetOperation_Child {
	return &v0.SetOperation_Child{
		ChildType: &v0.SetOperation_Child_TupleToUserset{
			TupleToUserset: &v0.TupleToUserset{
				Tupleset: &v0.TupleToUserset_Tupleset{
					Relation: tuplesetRelation,
				},
				ComputedUserset: &v0.ComputedUserset{
					Relation: usersetRelation,
					Object:   v0.ComputedUserset_TUPLE_USERSET_OBJECT,
				},
			},
		},
	}
}
