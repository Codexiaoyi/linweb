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

package linweb

import (
	"errors"
	"testing"

	"github.com/Codexiaoyi/linweb/test/mocks"

	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func TestCustomizePlugins(t *testing.T) {
	// Arrange:mock data
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// mock a request instance
	mock_request := mocks.NewMockIRequest(ctrl)
	mock_request.EXPECT().Param("mock_param").Return("test_param")
	// mock a new context instance
	mock_context := mocks.NewMockIContext(ctrl)
	mock_context.EXPECT().Request().Return(mock_request)
	mock_context.EXPECT().New().Return(mock_context)
	mock_context.EXPECT().Reset(gomock.Nil(), gomock.Nil(), gomock.Nil())
	// mock model
	mock_model := mocks.NewMockIModel(ctrl)
	mock_model.EXPECT().New(gomock.Nil()).Return(mock_model)
	mock_model.EXPECT().ModelError().Return(errors.New("test_mock_model error"))

	//Act
	linweb := NewLinWeb(ContextPlugin(mock_context), ModelPlugin(mock_model))
	ctx := linweb.markContext.New()
	ctx.Reset(nil, nil, nil)

	//Assert
	assert.Equal(t, "test_param", ctx.Request().Param("mock_param"))
	assert.Equal(t, "test_mock_model error", NewModel(nil).ModelError().Error())
}
