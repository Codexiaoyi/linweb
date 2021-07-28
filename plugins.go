package linweb

import (
	"linweb/interfaces"
	"linweb/pkg/model"
)

var plugins_model interfaces.IModel = &model.Model{}

// Create a new model plugin.
func NewModel(model interface{}) interfaces.IModel {
	return plugins_model.New(model)
}
