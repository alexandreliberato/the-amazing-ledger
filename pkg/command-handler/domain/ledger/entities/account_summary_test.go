package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountSummary(t *testing.T) {
	version := Version(1)
	accountPathLiability := "liability"
	accountPathAssets := "assets"

	paths := []Path{
		{
			*&accountPathLiability,
			200,
			300,
		},
		{
			*&accountPathAssets,
			400,
			500,
		},
	}

	accountSummary, err := NewAccountSummary(600, 800, paths, version)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(accountSummary.Paths))
}
