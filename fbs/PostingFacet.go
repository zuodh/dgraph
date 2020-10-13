package fbs

import "github.com/dgraph-io/dgraph/fb"

type postingFacetBuilder struct {
	p *PostingBuilder
	f *FacetBuilder
}

func newPostingFacetBuilder(p *PostingBuilder) *postingFacetBuilder {
	return &postingFacetBuilder{
		p: p,
		f: &FacetBuilder{
			builder: p.builder,
		},
	}
}

func (p *postingFacetBuilder) SetKey(key string) *postingFacetBuilder {
	p.f.SetKey(key)
	return p
}

func (p *postingFacetBuilder) SetValue(value []byte) *postingFacetBuilder {
	p.f.SetValue(value)
	return p
}

func (p *postingFacetBuilder) SetValueType(valueType fb.FacetValueType) *postingFacetBuilder {
	p.f.SetValueType(valueType)
	return p
}

func (p *postingFacetBuilder) SetTokens(tokens []string) *postingFacetBuilder {
	p.f.SetTokens(tokens)
	return p
}

func (p *postingFacetBuilder) SetAlias(alias string) *postingFacetBuilder {
	p.f.SetAlias(alias)
	return p
}

func (p *postingFacetBuilder) BuildFacet() *PostingBuilder {
	offset := p.f.buildOffset()
	p.p.facets = append(p.p.facets, offset)
	return p.p
}
