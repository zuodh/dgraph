package fbx_test

import (
	"fmt"
	"testing"

	"github.com/dgraph-io/dgraph/fb"
	"github.com/dgraph-io/dgraph/fbx"
	"github.com/dgraph-io/dgraph/protos/pb"
	"github.com/stretchr/testify/require"
)

func TestDirectedEdge(t *testing.T) {
	entity := uint64(1)
	attr := "attr"
	value := []byte("value")
	valueType := pb.Posting_BINARY
	valueID := uint64(2)
	label := "label"
	lang := "lang"
	op := fb.DirectedEdgeOpDEL
	facets := make([]*fb.Facet, 0)
	for i := 0; i < 5; i++ {
		facet := fbx.NewFacet().
			SetKey(fmt.Sprintf("facet%d", i)).
			Build()
		facets = append(facets, facet)
	}
	allowedPreds := []string{"some", "allowed", "preds"}

	builder := fbx.NewDirectedEdge().
		SetEntity(entity).
		SetAttr(attr).
		SetValue(value).
		SetValueType(valueType).
		SetValueID(valueID).
		SetLabel(label).
		SetLang(lang).
		SetOp(op).
		SetAllowedPreds(allowedPreds)

	for _, facet := range facets {
		builder.StartFacet().
			SetKey(fbx.BytesToString(facet.Key())).
			EndFacet()
	}

	de := builder.Build()

	require.Equal(t, de.Entity(), entity)
	require.Equal(t, fbx.BytesToString(de.Attr()), attr)
	require.Equal(t, de.ValueBytes(), value)
	require.Equal(t, pb.Posting_ValType(de.ValueType()), valueType)
	require.Equal(t, de.ValueId(), valueID)
	require.Equal(t, fbx.BytesToString(de.Label()), label)
	require.Equal(t, fbx.BytesToString(de.Lang()), lang)
	require.Equal(t, de.Op(), op)
	require.Equal(t, de.FacetsLength(), len(facets))
	for i, expFacet := range facets {
		var gotFacet fb.Facet
		require.True(t, de.Facets(&gotFacet, i))
		require.Equal(t, gotFacet.Key(), expFacet.Key())
	}
	require.Equal(t, de.AllowedPredsLength(), len(allowedPreds))
	for i, pred := range allowedPreds {
		require.Equal(t, fbx.BytesToString(de.AllowedPreds(i)), pred)
	}
}
