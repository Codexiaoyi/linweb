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

func TestAnnotateComment(t *testing.T) {
	// post,"aaa"
	comment := "[POST(\"aaa\")]"
	resType, resString := annotateComment(comment)
	assert.Equal(t, POST, resType)
	assert.Equal(t, "aaa", resString)

	// unknown MethodType, Result => unknown,"aaa"
	comment = "[POS(\"aaa\")]"
	resType, resString = annotateComment(comment)
	assert.Equal(t, Unknown, resType)
	assert.Equal(t, "aaa", resString)

	// not start with '[', Result => post,"aaa"
	comment = "#[POST(\"aaa\")]"
	resType, resString = annotateComment(comment)
	assert.Equal(t, POST, resType)
	assert.Equal(t, "aaa", resString)

	// not exist '(' or ')', Result => unknown,""
	comment = "[POST\"aaa\"]"
	resType, resString = annotateComment(comment)
	assert.Equal(t, Unknown, resType)
	assert.Equal(t, "", resString)

	// not exist '[' or ']', Result => unknown,""
	comment = "POST(\"aaa\")"
	resType, resString = annotateComment(comment)
	assert.Equal(t, Unknown, resType)
	assert.Equal(t, "", resString)
}
