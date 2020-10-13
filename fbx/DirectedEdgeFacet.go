package fbx

import (
	"github.com/dgraph-io/dgo/v200/protos/api"
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

func (f *directedEdgeFacet) From(facet *api.Facet) *directedEdgeFacet {
	f.f.From(facet)
	return f
}

func (f *directedEdgeFacet) SetKey(key string) *directedEdgeFacet {
	f.f.SetKey(key)
	return f
}

func (f *directedEdgeFacet) SetValue(value []byte) *directedEdgeFacet {
	f.f.SetValue(value)
	return f
}

func (f *directedEdgeFacet) SetValueType(valueType api.Facet_ValType) *directedEdgeFacet {
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

func (f *directedEdgeFacet) EndFacet() *DirectedEdge {
	offset := f.f.buildOffset()
	f.de.facets = append(f.de.facets, offset)
	return f.de
}
