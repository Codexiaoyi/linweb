//Copyright 2021 Codexiaoyi
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

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

	gp_ok := []string{"get", "linweb", "myfile", "ok"}
	assert.Equal(t, "/get/:name/myfile/ok", node.search(gp_ok, 0).url)

	gp_any := []string{"get", "linweb", "myfile", "any"}
	assert.Equal(t, "/get/:name/*filepath/ok", node.search(gp_any, 0).url)

	// "/get/:name/*filepath/aa" cover "/get/:name/*filepath/ok"
	node.insert("/get/:name/*filepath/aa", parts, 0)
	assert.Equal(t, "/get/:name/*filepath/aa", node.search(gp_any, 0).url)
}
