package linweb

import (
	"errors"
	"linweb/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCustomizePlugins(t *testing.T) {
	linweb := NewLinweb()
	linweb.AddCustomizePlugins(&demoModel{})
	assert.Equal(t, "TestModelDemo", NewModel(nil).Validate().Error())
}

type demoModel struct {
	m interface{}
}

func (m *demoModel) New(model interface{}) interfaces.IModel {
	return &demoModel{
		m: model,
	}
}

func (m *demoModel) Validate() error {
	return errors.New("TestModelDemo")
}

func (m *demoModel) MapToByFieldName(dest interface{}) error {

	return nil
}

func (m *demoModel) MapToByFieldTag(dest interface{}) error {
	//TODO
	return nil
}

func (m *demoModel) MapFromByFieldName(src interface{}) error {

	return nil
}

func (m *demoModel) MapFromByFieldTag(src interface{}) error {
	//TODO
	return nil
}
