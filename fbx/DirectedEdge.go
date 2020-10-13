package fbx

import (
	"github.com/dgraph-io/dgraph/fb"
	flatbuffers "github.com/google/flatbuffers/go"
)

type DirectedEdge struct {
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

func NewDirectedEdge() *DirectedEdge {
	return &DirectedEdge{
		builder: flatbuffers.NewBuilder(bufSize),
	}
}

func (de *DirectedEdge) SetEntity(entity uint64) *DirectedEdge {
	de.entity = entity
	return de
}

func (de *DirectedEdge) SetAttr(attr string) *DirectedEdge {
	de.attr = de.builder.CreateString(attr)
	return de
}

func (de *DirectedEdge) SetValue(value []byte) *DirectedEdge {
	de.value = de.builder.CreateByteVector(value)
	return de
}

func (de *DirectedEdge) SetValueType(valueType fb.PostingValueType) *DirectedEdge {
	de.valueType = valueType
	return de
}

func (de *DirectedEdge) SetValueID(valueID uint64) *DirectedEdge {
	de.valueID = valueID
	return de
}

func (de *DirectedEdge) SetLabel(label string) *DirectedEdge {
	de.label = de.builder.CreateString(label)
	return de
}

func (de *DirectedEdge) SetLang(lang string) *DirectedEdge {
	de.lang = de.builder.CreateString(lang)
	return de
}

func (de *DirectedEdge) SetOp(op fb.DirectedEdgeOp) *DirectedEdge {
	de.op = op
	return de
}

func (de *DirectedEdge) AppendFacet() *directedEdgeFacet {
	return newDirectedEdgeFacet(de)
}

func (de *DirectedEdge) SetAllowedPreds(allowedPreds []string) *DirectedEdge {
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

func (de *DirectedEdge) buildOffset() flatbuffers.UOffsetT {
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

func (de *DirectedEdge) Build() *fb.DirectedEdge {
	directedEdge := de.buildOffset()
	de.builder.Finish(directedEdge)
	buf := de.builder.FinishedBytes()
	return fb.GetRootAsDirectedEdge(buf, 0)
}
