package linweb

import (
	"errors"
	"testing"

	"linweb/test/mocks"

	"github.com/golang/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func TestAddCustomizePlugins(t *testing.T) {
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
	linweb := NewLinWeb()
	linweb.AddCustomizePlugins(mock_context, mock_model)
	ctx := linweb.markContext.New()
	ctx.Reset(nil, nil, nil)

	//Assert
	assert.Equal(t, "test_param", ctx.Request().Param("mock_param"))
	assert.Equal(t, "test_mock_model error", NewModel(nil).ModelError().Error())
}
