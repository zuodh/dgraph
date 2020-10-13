package fbx

import (
	"github.com/dgraph-io/dgraph/fb"
)

type directedEdgeFacet struct {
	de *DirectedEdge
	f  *Facet
}

func newDirectedEdgeFacet(de *DirectedEdge) *directedEdgeFacet {
	return &directedEdgeFacet{
		de: de,
		f: &Facet{
			builder: de.builder,
		},
	}
}

func (f *directedEdgeFacet) SetKey(key string) *directedEdgeFacet {
	f.f.SetKey(key)
	return f
}

func (f *directedEdgeFacet) SetValue(value []byte) *directedEdgeFacet {
	f.f.SetValue(value)
	return f
}

func (f *directedEdgeFacet) SetValueType(valueType fb.FacetValueType) *directedEdgeFacet {
	f.f.SetValueType(valueType)
	return f
}

func (f *directedEdgeFacet) SetTokens(tokens []string) *directedEdgeFacet {
	f.f.SetTokens(tokens)
	return f
}

func (f *directedEdgeFacet) SetAlias(alias string) *directedEdgeFacet {
	f.f.SetAlias(alias)
	return f
}

func (f *directedEdgeFacet) BuildFacet() *DirectedEdge {
	offset := f.f.buildOffset()
	f.de.facets = append(f.de.facets, offset)
	return f.de
}
