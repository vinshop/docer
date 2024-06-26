package docer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRead(t *testing.T) {
	doc, err := Read("test.json")
	assert.NoError(t, err)
	assert.NotNil(t, doc)
}

func TestDoc_Generate(t *testing.T) {
	doc, err := Read("test.json")
	assert.NoError(t, err)
	assert.NotNil(t, doc)

	err = doc.Generate("test.md")
	assert.NoError(t, err)
}
