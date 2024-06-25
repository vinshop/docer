package doc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRead(t *testing.T) {
	doc, err := Read("test.json")
	assert.NoError(t, err)
	assert.NotNil(t, doc)
}
