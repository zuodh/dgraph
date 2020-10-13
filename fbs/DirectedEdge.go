package fbs

import (
	"github.com/dgraph-io/dgraph/fb"
	flatbuffers "github.com/google/flatbuffers/go"
)

type DirectedEdgeBuilder struct {
	builder *flatbuffers.Builder

	entity       uint64
	attr         flatbuffers.UOffsetT
	value        flatbuffers.UOffsetT
	valueType    fb.PostingValueType
	valueID      uint64
	label        flatbuffers.UOffsetT
	lang         flatbuffers.UOffsetT
	op           fb.DirectedEdgeOp
	facets       []flatbuffers.UOffsetT
	allowedPreds flatbuffers.UOffsetT
}

func NewDirectedEdgeBuilder() *DirectedEdgeBuilder {
	return &DirectedEdgeBuilder{
		builder: flatbuffers.NewBuilder(bufSize),
	}
}

func (de *DirectedEdgeBuilder) SetEntity(entity uint64) *DirectedEdgeBuilder {
	de.entity = entity
	return de
}

func (de *DirectedEdgeBuilder) SetAttr(attr string) *DirectedEdgeBuilder {
	de.attr = de.builder.CreateString(attr)
	return de
}

func (de *DirectedEdgeBuilder) SetValue(value []byte) *DirectedEdgeBuilder {
	de.value = de.builder.CreateByteVector(value)
	return de
}

func (de *DirectedEdgeBuilder) SetValueType(valueType fb.PostingValueType) *DirectedEdgeBuilder {
	de.valueType = valueType
	return de
}

func (de *DirectedEdgeBuilder) SetValueID(valueID uint64) *DirectedEdgeBuilder {
	de.valueID = valueID
	return de
}

func (de *DirectedEdgeBuilder) SetLabel(label string) *DirectedEdgeBuilder {
	de.label = de.builder.CreateString(label)
	return de
}

func (de *DirectedEdgeBuilder) SetLang(lang string) *DirectedEdgeBuilder {
	de.lang = de.builder.CreateString(lang)
	return de
}

func (de *DirectedEdgeBuilder) SetOp(op fb.DirectedEdgeOp) *DirectedEdgeBuilder {
	de.op = op
	return de
}

func (de *DirectedEdgeBuilder) AppendFacet() *directedEdgeFacetBuilder {
	return newDirectedEdgeFacetBuilder(de)
}

func (de *DirectedEdgeBuilder) SetAllowedPreds(allowedPreds []string) *DirectedEdgeBuilder {
	offsets := make([]flatbuffers.UOffsetT, len(allowedPreds))
	for i, pred := range allowedPreds {
		offsets[i] = de.builder.CreateString(pred)
	}

	fb.FacetStartTokensVector(de.builder, len(allowedPreds))
	for i := len(allowedPreds) - 1; i >= 0; i-- {
		de.builder.PrependUOffsetT(offsets[i])
	}
	de.allowedPreds = de.builder.EndVector(len(allowedPreds))

	return de
}

func (de *DirectedEdgeBuilder) buildOffset() flatbuffers.UOffsetT {
	fb.DirectedEdgeStartFacetsVector(de.builder, len(de.facets))
	for i := len(de.facets) - 1; i >= 0; i-- {
		de.builder.PrependUOffsetT(de.facets[i])
	}
	facets := de.builder.EndVector(len(de.facets))

	fb.DirectedEdgeStart(de.builder)
	fb.DirectedEdgeAddEntity(de.builder, de.entity)
	fb.DirectedEdgeAddAttr(de.builder, de.attr)
	fb.DirectedEdgeAddValue(de.builder, de.value)
	fb.DirectedEdgeAddValueType(de.builder, de.valueType)
	fb.DirectedEdgeAddValueId(de.builder, de.valueID)
	fb.DirectedEdgeAddLabel(de.builder, de.label)
	fb.DirectedEdgeAddLang(de.builder, de.lang)
	fb.DirectedEdgeAddOp(de.builder, de.op)
	fb.DirectedEdgeAddFacets(de.builder, facets)
	fb.DirectedEdgeAddAllowedPreds(de.builder, de.allowedPreds)
	return fb.DirectedEdgeEnd(de.builder)
}

func (de *DirectedEdgeBuilder) Build() *fb.DirectedEdge {
	directedEdge := de.buildOffset()
	de.builder.Finish(directedEdge)
	buf := de.builder.FinishedBytes()
	return fb.GetRootAsDirectedEdge(buf, 0)
}
