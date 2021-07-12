package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	node := &node{}
	parts := []string{"get", "linweb", ":id"}
	node.insert("/get/linweb/:id", parts, 0)
	for i := 0; i < len(parts); i++ {
		for _, c := range node.children {
			if c.part == parts[i] {
				node = c
				break
			}
		}
		assert.Equal(t, parts[i], node.part)
	}

}

func TestInsert1(t *testing.T) {
	node := &node{}
	parts1 := []string{"get", ":name", "*filepath", "ok"}
	node.insert("/get/:name/*filepath/ok", parts1, 0)
	assert.Equal(t, "get", node.children[0].part)
	assert.Equal(t, ":name", node.children[0].children[0].part)
	assert.Equal(t, "*filepath", node.children[0].children[0].children[0].part)
	assert.Equal(t, "/get/:name/*filepath/ok", node.children[0].children[0].children[0].url)
	assert.Equal(t, 0, len(node.children[0].children[0].children[0].children))
}

func TestSearch(t *testing.T) {
	node := &node{}
	parts := []string{"get", ":name", "*filepath", "ok"}
	parts1 := []string{"get", ":name", "myfile", "ok"}
	node.insert("/get/:name/myfile/ok", parts1, 0)
	node.insert("/get/:name/*filepath/ok", parts, 0)

	gp := []string{"get", "linweb", "myfile", "aa"}
	res := node.search(gp, 0)
	assert.Equal(t, "/get/:name/*filepath/ok", res.url)
}
