package router

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestRouter() *Router {
	r := NewRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	assert.Equal(t, []string{"p", ":name"}, parsePattern("/p/:name"))
	assert.Equal(t, []string{"p", "*"}, parsePattern("/p/*"))
	assert.Equal(t, []string{"p", "*name"}, parsePattern("/p/*name/*"))
}

func TestGetRoute(t *testing.T) {
	router := newTestRouter()
	node, params := router.getRoute("GET", "/hello/linweb")
	assert.NotEmpty(t, node)
	assert.Equal(t, "/hello/:name", node.url)
	assert.Equal(t, "linweb", params["name"])

	node1, params1 := router.getRoute("GET", "/assets/go/linweb/router.go")
	assert.NotEmpty(t, node1)
	assert.Equal(t, "/assets/*filepath", node1.url)
	assert.Equal(t, "go/linweb/router.go", params1["filepath"])
}
