package fbs_test

import (
	"fmt"
	"testing"

	"github.com/dgraph-io/dgraph/fb"
	"github.com/dgraph-io/dgraph/fbs"
	"github.com/stretchr/testify/require"
)

func TestDirectedEdgeBuilder(t *testing.T) {
	entity := uint64(1)
	attr := "attr"
	value := []byte("value")
	valueType := fb.PostingValueTypeBINARY
	valueID := uint64(2)
	label := "label"
	lang := "lang"
	op := fb.DirectedEdgeOpDEL
	facets := make([]*fb.Facet, 0)
	for i := 0; i < 5; i++ {
		facet := fbs.NewFacetBuilder().
			SetKey(fmt.Sprintf("facet%d", i)).
			Build()
		facets = append(facets, facet)
	}
	allowedPreds := []string{"some", "allowed", "preds"}

	builder := fbs.NewDirectedEdgeBuilder().
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
		builder.AppendFacet().
			SetKey(fbs.BytesToString(facet.Key())).
			BuildFacet()
	}

	de := builder.Build()

	require.Equal(t, de.Entity(), entity)
	require.Equal(t, fbs.BytesToString(de.Attr()), attr)
	require.Equal(t, de.ValueBytes(), value)
	require.Equal(t, de.ValueType(), valueType)
	require.Equal(t, de.ValueId(), valueID)
	require.Equal(t, fbs.BytesToString(de.Label()), label)
	require.Equal(t, fbs.BytesToString(de.Lang()), lang)
	require.Equal(t, de.Op(), op)
	require.Equal(t, de.FacetsLength(), len(facets))
	for i, expFacet := range facets {
		var gotFacet fb.Facet
		require.True(t, de.Facets(&gotFacet, i))
		require.Equal(t, gotFacet.Key(), expFacet.Key())
	}
	require.Equal(t, de.AllowedPredsLength(), len(allowedPreds))
	for i, pred := range allowedPreds {
		require.Equal(t, fbs.BytesToString(de.AllowedPreds(i)), pred)
	}
}
