package gsheet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	links, err := convertLinksDto(``)
	assert.Nil(t, err)
	assert.Empty(t, links)

	links, err = convertLinksDto(`
	Первый день: http://google.com

	`)
	assert.Nil(t, err)
	assert.Len(t, links, 1)
	assert.Equal(t, TextLink{Text: "Первый день", Link: "http://google.com"}, links[0])
}
