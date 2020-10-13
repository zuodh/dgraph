package fbs

import (
	"github.com/dgraph-io/dgraph/fb"
	flatbuffers "github.com/google/flatbuffers/go"
)

type FacetBuilder struct {
	builder *flatbuffers.Builder

	key       flatbuffers.UOffsetT
	value     flatbuffers.UOffsetT
	valueType fb.FacetValueType
	tokens    flatbuffers.UOffsetT
	alias     flatbuffers.UOffsetT
}

func NewFacetBuilder() *FacetBuilder {
	return &FacetBuilder{
		builder: flatbuffers.NewBuilder(bufSize),
	}
}

func (f *FacetBuilder) SetKey(key string) *FacetBuilder {
	f.key = f.builder.CreateString(key)
	return f
}

func (f *FacetBuilder) SetValue(value []byte) *FacetBuilder {
	f.value = f.builder.CreateByteVector(value)
	return f
}

func (f *FacetBuilder) SetValueType(valueType fb.FacetValueType) *FacetBuilder {
	f.valueType = valueType
	return f
}

func (f *FacetBuilder) SetTokens(tokens []string) *FacetBuilder {
	offsets := make([]flatbuffers.UOffsetT, len(tokens))
	for i, token := range tokens {
		offsets[i] = f.builder.CreateString(token)
	}

	fb.FacetStartTokensVector(f.builder, len(tokens))
	for i := len(tokens) - 1; i >= 0; i-- {
		f.builder.PrependUOffsetT(offsets[i])
	}
	f.tokens = f.builder.EndVector(len(tokens))

	return f
}

func (f *FacetBuilder) SetAlias(alias string) *FacetBuilder {
	f.alias = f.builder.CreateString(alias)
	return f
}

func (f *FacetBuilder) Build() *fb.Facet {
	facet := f.buildOffset()
	f.builder.Finish(facet)

	buf := f.builder.FinishedBytes()
	return fb.GetRootAsFacet(buf, 0)
}

func (f *FacetBuilder) buildOffset() flatbuffers.UOffsetT {
	fb.FacetStart(f.builder)
	fb.FacetAddKey(f.builder, f.key)
	fb.FacetAddValue(f.builder, f.value)
	fb.FacetAddValueType(f.builder, f.valueType)
	fb.FacetAddTokens(f.builder, f.tokens)
	fb.FacetAddAlias(f.builder, f.alias)
	return fb.FacetEnd(f.builder)
}
