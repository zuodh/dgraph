package fbx

import (
	"github.com/dgraph-io/dgraph/fb"
	flatbuffers "github.com/google/flatbuffers/go"
)

type Posting struct {
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

func NewPosting() *Posting {
	return &Posting{
		builder: flatbuffers.NewBuilder(bufSize),
	}
}

func (p *Posting) SetUid(uid uint64) *Posting {
	p.uid = uid
	return p
}

func (p *Posting) SetValue(value []byte) *Posting {
	p.value = p.builder.CreateByteVector(value)
	return p
}

func (p *Posting) SetValueType(valueType fb.PostingValueType) *Posting {
	p.valueType = valueType
	return p
}

func (p *Posting) SetLangTag(langTag []byte) *Posting {
	p.langTag = p.builder.CreateByteVector(langTag)
	return p
}

func (p *Posting) SetLabel(label string) *Posting {
	p.label = p.builder.CreateString(label)
	return p
}

func (p *Posting) AppendFacet() *postingFacet {
	return newPostingFacet(p)
}

func (p *Posting) SetOp(op fb.DirectedEdgeOp) *Posting {
	p.op = op
	return p
}

func (p *Posting) SetStartTs(startTs uint64) *Posting {
	p.startTs = startTs
	return p
}

func (p *Posting) SetCommitTs(commitTs uint64) *Posting {
	p.commitTs = commitTs
	return p
}

func (p *Posting) buildOffset() flatbuffers.UOffsetT {
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

func (p *Posting) Build() *fb.Posting {
	posting := p.buildOffset()
	p.builder.Finish(posting)
	buf := p.builder.FinishedBytes()
	return fb.GetRootAsPosting(buf, 0)
}
