package docer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNew(t *testing.T) {

	// string
	assert.Equal(t, "old", getNew("old", ""))
	assert.Equal(t, "new", getNew("old", "new"))
	// int
	assert.Equal(t, 1, getNew(1, 0))
	assert.Equal(t, 1, getNew(1, 1))
	assert.Equal(t, 2, getNew(1, 2))
	// arr
	assert.Equal(t, []int{1}, getNew([]int{1}, []int{}))
	assert.Equal(t, []int{2}, getNew([]int{1, 2, 3}, []int{2}))
	// nil
	var old any
	old = []int{1}
	assert.Equal(t, old, getNew(old, nil))
}

func TestMergeField(t *testing.T) {
	old := &Field{
		Name:        "old",
		Type:        "old",
		Required:    true,
		Ref:         "old",
		Description: "old",
	}
	new := &Field{
		Name:        "new",
		Type:        "",
		Required:    false,
		Ref:         "new",
		Description: "",
	}
	res := mergeField(old, new)
	assert.Equal(t, "new", res.Name)
	assert.Equal(t, "old", res.Type)
	assert.Equal(t, true, res.Required)
	assert.Equal(t, "new", res.Ref)
	assert.Equal(t, "old", res.Description)
}

func TestMergeType(t *testing.T) {
	old := &Type{
		Name:        "A",
		DisplayName: "old",
		Description: "old",
		Fields: []*Field{
			{
				Name:        "A",
				Type:        "1",
				Required:    false,
				Ref:         "2",
				Description: "3",
			},
			{
				Name:        "B",
				Type:        "4",
				Required:    true,
				Ref:         "5",
				Description: "6",
			},
		},
	}
	new := &Type{
		Name:        "A",
		DisplayName: "",
		Description: "",
		Fields: []*Field{
			{
				Name:        "B",
				Type:        "7",
				Required:    false,
				Ref:         "",
				Description: "",
			},
		},
	}
	res := mergeType(old, new)
	assert.Equal(t, "A", res.Name)
	assert.Equal(t, "old", res.DisplayName)
	assert.Equal(t, "old", res.Description)
	assert.Equal(t, 1, len(res.Fields))
	assert.Equal(t, "B", res.Fields[0].Name)
	assert.Equal(t, "7", res.Fields[0].Type)
	assert.Equal(t, true, res.Fields[0].Required)
	assert.Equal(t, "5", res.Fields[0].Ref)
	assert.Equal(t, "6", res.Fields[0].Description)
}

func TestMergeDoc(t *testing.T) {
	old := &Doc{
		URL:             "old",
		Method:          "old",
		Headers:         []string{"old"},
		ExampleBody:     "old",
		SuccessResponse: "old",
		ErrorResponse:   "old",
		Types: []*Type{
			{},
			{
				Name:        "A",
				DisplayName: "old",
				Description: "old",
				Fields: []*Field{
					{
						Name:        "A",
						Type:        "1",
						Required:    false,
						Ref:         "2",
						Description: "3",
					},
					{
						Name:        "B",
						Type:        "4",
						Required:    true,
						Ref:         "5",
						Description: "6",
					},
				},
			},
		},
	}
	new := &Doc{
		URL:             "new",
		Method:          "",
		Headers:         nil,
		ExampleBody:     nil,
		SuccessResponse: nil,
		ErrorResponse:   nil,
		Types: []*Type{
			{
				Name:        "A",
				DisplayName: "",
				Description: "",
				Fields: []*Field{
					{
						Name:        "B",
						Type:        "7",
						Required:    false,
						Ref:         "",
						Description: "",
					},
				},
			},
		},
	}
	res := mergeDoc(old, new)
	assert.Equal(t, "new", res.URL)
	assert.Equal(t, "old", res.Method)
	assert.Equal(t, 1, len(res.Headers))
	assert.Equal(t, "old", res.Headers[0])
	assert.Equal(t, "old", res.ExampleBody)
	assert.Equal(t, "old", res.SuccessResponse)
	assert.Equal(t, "old", res.ErrorResponse)
	assert.Equal(t, 1, len(res.Types))
	assert.Equal(t, "A", res.Types[0].Name)
	assert.Equal(t, "old", res.Types[0].DisplayName)
	assert.Equal(t, "old", res.Types[0].Description)
	assert.Equal(t, 1, len(res.Types[0].Fields))
	assert.Equal(t, "B", res.Types[0].Fields[0].Name)
	assert.Equal(t, "7", res.Types[0].Fields[0].Type)
	assert.Equal(t, true, res.Types[0].Fields[0].Required)
	assert.Equal(t, "5", res.Types[0].Fields[0].Ref)
	assert.Equal(t, "6", res.Types[0].Fields[0].Description)

}
