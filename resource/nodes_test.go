package resource

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/containership/csctl/cloud/provision/types"
)

var (
	nodeTime = "1517001176920"

	nodes = []types.Node{
		{
			ID: types.UUID("1234"),
			Status: &types.NodeStatus{
				Type: strptr("RUNNING"),
			},
			CreatedAt: &nodeTime,
			UpdatedAt: &nodeTime,
		},
		{
			ID: types.UUID("4321"),
			Status: &types.NodeStatus{
				Type: strptr("CREATING"),
			},
			CreatedAt: &nodeTime,
			UpdatedAt: &nodeTime,
		},
	}

	nodesSingle = []types.Node{
		{
			ID: types.UUID("1234"),
			Status: &types.NodeStatus{
				Type: strptr("RUNNING"),
			},
			CreatedAt: &nodeTime,
			UpdatedAt: &nodeTime,
		},
	}
)

func TestNodesJSON(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewNodes(nodesSingle)
	err := a.JSON(buf, true)
	assert.Nil(t, err)

	err = a.JSON(buf, false)
	assert.Nil(t, err)
}

func TestNodesYAML(t *testing.T) {
	buf := new(bytes.Buffer)
	a := NewNodes(nodesSingle)
	err := a.YAML(buf, true)
	assert.Nil(t, err)

	err = a.YAML(buf, false)
	assert.Nil(t, err)
}

func TestNewNodes(t *testing.T) {
	a := NewNodes(nil)
	assert.NotNil(t, a)

	a = NewNodes(nodes)
	assert.NotNil(t, a)
	assert.Equal(t, len(a.items), len(nodes))

	a = Node()
	assert.NotNil(t, a)
}

func TestNodesTable(t *testing.T) {
	buf := new(bytes.Buffer)

	a := NewNodes(nodes)
	assert.NotNil(t, a)

	err := a.Table(buf)
	assert.Nil(t, err)

	info, err := getTableInfo(buf)
	assert.Nil(t, err)
	assert.Equal(t, len(a.columns()), info.numHeaderCols)
	assert.Equal(t, len(a.columns()), info.numCols)
	assert.Equal(t, len(nodes), info.numRows)
}
