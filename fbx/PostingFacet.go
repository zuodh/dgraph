package fbx

import (
	"github.com/dgraph-io/dgo/v200/protos/api"
)

type postingFacet struct {
	p *Posting
	f *Facet
}

func newPostingFacet(p *Posting) *postingFacet {
	return &postingFacet{
		p: p,
		f: &Facet{
			builder: p.builder,
		},
	}
}

func (p *postingFacet) SetKey(key string) *postingFacet {
	p.f.SetKey(key)
	return p
}

func (p *postingFacet) SetValue(value []byte) *postingFacet {
	p.f.SetValue(value)
	return p
}

func (p *postingFacet) SetValueType(valueType api.Facet_ValType) *postingFacet {
	p.f.SetValueType(valueType)
	return p
}

func (p *postingFacet) SetTokens(tokens []string) *postingFacet {
	p.f.SetTokens(tokens)
	return p
}

func (p *postingFacet) SetAlias(alias string) *postingFacet {
	p.f.SetAlias(alias)
	return p
}

func (p *postingFacet) BuildFacet() *Posting {
	offset := p.f.buildOffset()
	p.p.facets = append(p.p.facets, offset)
	return p.p
}
