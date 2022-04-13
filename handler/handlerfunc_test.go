package handler

import (
	"github.com/stretchr/testify/require"
	"new_ozon_test/storage"
	"sync"
	"testing"
)

func TestApp_PasteAndGetLink(t *testing.T) {

	fullLink := "test := handler.App{}test.StorType = &Memory{}"
	test := App{}
	output := storage.Data{}
	test.StorType = &storage.Memory{
		LongLinks:  make(map[string]string),
		ShortLinks: make(map[string]string),
		Mu:         new(sync.Mutex),
	}
	output, err := test.StorType.AddLink(fullLink)
	require.NoError(t, err, "should not have err")
	require.NotEmpty(t, output.ShortLink, "")
	output, err = test.StorType.GetLink(output.ShortLink)
	require.NoError(t, err, "should not have err")
	require.Equal(t, fullLink, output.FullLink, "should by eql")

}
