package fbx_test

import (
	"testing"

	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/dgraph-io/dgraph/fbx"
	"github.com/stretchr/testify/require"
)

func TestFacet(t *testing.T) {
	key := "key"
	value := []byte("value")
	valueType := api.Facet_STRING
	tokens := []string{"some", "tokens"}
	alias := "alias"

	facet := fbx.NewFacet().
		SetKey(key).
		SetValue(value).
		SetValueType(valueType).
		SetTokens(tokens).
		SetAlias(alias).
		Build()

	require.Equal(t, fbx.BytesToString(facet.Key()), key)
	require.Equal(t, facet.ValueBytes(), value)
	require.Equal(t, facet.TokensLength(), len(tokens))
	for i, token := range tokens {
		require.Equal(t, fbx.BytesToString(facet.Tokens(i)), token)
	}
	require.Equal(t, fbx.BytesToString(facet.Alias()), alias)
}
