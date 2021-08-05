package model

import (
	"linweb/interfaces"

	"github.com/Codexiaoyi/go-mapper"
	"github.com/go-playground/validator"
)

type Model struct {
	model interface{}
	err   error
}

func (m *Model) New(model interface{}) interfaces.IModel {
	return &Model{
		model: model,
		err:   nil,
	}
}

func (model *Model) Validate() interfaces.IModel {
	if model.err != nil {
		return model
	}
	v := validator.New()
	err := v.Struct(model.model)
	if err != nil {
		model.err = err
	}
	return model
}

func (model *Model) MapToByFieldName(dest interface{}) interfaces.IModel {
	if model.err != nil {
		return model
	}
	err := mapper.StructMapByFieldName(model.model, dest)
	if err != nil {
		model.err = err
	}
	return model
}

func (model *Model) MapToByFieldTag(dest interface{}) interfaces.IModel {
	if model.err != nil {
		return model
	}
	err := mapper.StructMapByFieldTag(model.model, dest)
	if err != nil {
		model.err = err
	}
	return model
}

func (model *Model) MapFromByFieldName(src interface{}) interfaces.IModel {
	if model.err != nil {
		return model
	}
	err := mapper.StructMapByFieldName(src, model.model)
	if err != nil {
		model.err = err
	}
	return model
}

func (model *Model) MapFromByFieldTag(src interface{}) interfaces.IModel {
	if model.err != nil {
		return model
	}
	err := mapper.StructMapByFieldTag(src, model.model)
	if err != nil {
		model.err = err
	}
	return model
}

func (model *Model) ModelError() error {
	return model.err
}
