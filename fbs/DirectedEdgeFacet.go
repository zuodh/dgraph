package fbs

import (
	"github.com/dgraph-io/dgraph/fb"
)

type directedEdgeFacetBuilder struct {
	de *DirectedEdgeBuilder
	f  *FacetBuilder
}

func newDirectedEdgeFacetBuilder(de *DirectedEdgeBuilder) *directedEdgeFacetBuilder {
	return &directedEdgeFacetBuilder{
		de: de,
		f: &FacetBuilder{
			builder: de.builder,
		},
	}
}

func (f *directedEdgeFacetBuilder) SetKey(key string) *directedEdgeFacetBuilder {
	f.f.SetKey(key)
	return f
}

func (f *directedEdgeFacetBuilder) SetValue(value []byte) *directedEdgeFacetBuilder {
	f.f.SetValue(value)
	return f
}

func (f *directedEdgeFacetBuilder) SetValueType(valueType fb.FacetValueType) *directedEdgeFacetBuilder {
	f.f.SetValueType(valueType)
	return f
}

func (f *directedEdgeFacetBuilder) SetTokens(tokens []string) *directedEdgeFacetBuilder {
	f.f.SetTokens(tokens)
	return f
}

func (f *directedEdgeFacetBuilder) SetAlias(alias string) *directedEdgeFacetBuilder {
	f.f.SetAlias(alias)
	return f
}

func (f *directedEdgeFacetBuilder) BuildFacet() *DirectedEdgeBuilder {
	offset := f.f.buildOffset()
	f.de.facets = append(f.de.facets, offset)
	return f.de
}
