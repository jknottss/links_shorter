package createlink

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateLink(t *testing.T) {

	allowChars := "[0-9a-zA-Z_]{10}"

	for i := 0; i < 100; i++ {
		link := CreateLink()
		require.Regexp(t, allowChars, link, "contains allowed chars")
		require.NotEqualf(t, link, CreateLink(), "should not eql")
	}
}
