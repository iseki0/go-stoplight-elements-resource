package stoplightelementsres

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFS(t *testing.T) {
	var e error
	_, e = FS.Open("web-components.min.js")
	assert.NoError(t, e)
	_, e = FS.Open("style.min.css")
	assert.NoError(t, e)
}
