package fbs

import (
	"github.com/dgraph-io/dgraph/fb"
	flatbuffers "github.com/google/flatbuffers/go"
)

type PostingBuilder struct {
	builder *flatbuffers.Builder

	uid       uint64
	value     flatbuffers.UOffsetT
	valueType fb.PostingValueType
	langTag   flatbuffers.UOffsetT
	label     flatbuffers.UOffsetT
	facets    []flatbuffers.UOffsetT
	op        fb.DirectedEdgeOp
	startTs   uint64
	commitTs  uint64
}

func NewPostingBuilder() *PostingBuilder {
	return &PostingBuilder{
		builder: flatbuffers.NewBuilder(bufSize),
	}
}

func (p *PostingBuilder) SetUid(uid uint64) *PostingBuilder {
	p.uid = uid
	return p
}

func (p *PostingBuilder) SetValue(value []byte) *PostingBuilder {
	p.value = p.builder.CreateByteVector(value)
	return p
}

func (p *PostingBuilder) SetValueType(valueType fb.PostingValueType) *PostingBuilder {
	p.valueType = valueType
	return p
}

func (p *PostingBuilder) SetLangTag(langTag []byte) *PostingBuilder {
	p.langTag = p.builder.CreateByteVector(langTag)
	return p
}

func (p *PostingBuilder) SetLabel(label string) *PostingBuilder {
	p.label = p.builder.CreateString(label)
	return p
}

func (p *PostingBuilder) AppendFacet() *postingFacetBuilder {
	return newPostingFacetBuilder(p)
}

func (p *PostingBuilder) SetOp(op fb.DirectedEdgeOp) *PostingBuilder {
	p.op = op
	return p
}

func (p *PostingBuilder) SetStartTs(startTs uint64) *PostingBuilder {
	p.startTs = startTs
	return p
}

func (p *PostingBuilder) SetCommitTs(commitTs uint64) *PostingBuilder {
	p.commitTs = commitTs
	return p
}

func (p *PostingBuilder) buildOffset() flatbuffers.UOffsetT {
	fb.PostingStartFacetsVector(p.builder, len(p.facets))
	for i := len(p.facets) - 1; i >= 0; i-- {
		p.builder.PrependUOffsetT(p.facets[i])
	}
	facets := p.builder.EndVector(len(p.facets))

	fb.PostingStart(p.builder)
	fb.PostingAddUid(p.builder, p.uid)
	fb.PostingAddValue(p.builder, p.value)
	fb.PostingAddValueType(p.builder, p.valueType)
	fb.PostingAddLangTag(p.builder, p.langTag)
	fb.PostingAddLabel(p.builder, p.label)
	fb.PostingAddFacets(p.builder, facets)
	fb.PostingAddOp(p.builder, p.op)
	fb.PostingAddStartTs(p.builder, p.startTs)
	fb.PostingAddCommitTs(p.builder, p.commitTs)
	return fb.PostingEnd(p.builder)
}

func (p *PostingBuilder) Build() *fb.Posting {
	posting := p.buildOffset()
	p.builder.Finish(posting)
	buf := p.builder.FinishedBytes()
	return fb.GetRootAsPosting(buf, 0)
}
