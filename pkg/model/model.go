package model

import (
	"linweb/interfaces"

	"github.com/Codexiaoyi/go-mapper"
	"github.com/go-playground/validator"
)

type Model struct {
	m   interface{}
	err error
}

func New(model interface{}) interfaces.IModel {
	return &Model{
		m: model,
	}
}

func (m *Model) Validate() interfaces.IModel {
	if m.err != nil {
		return m
	}

	v := validator.New()
	err := v.Struct(m.m)
	if err != nil {
		m.err = err
	}

	return m
}

func (m *Model) MapToByFieldName(dest interface{}) interfaces.IModel {
	if m.err != nil {
		return m
	}
	err := mapper.StructMapByFieldName(m.m, dest)
	if err != nil {
		m.err = err
	}
	return m
}

func (m *Model) MapToByFieldTag(dest interface{}) interfaces.IModel {
	//TODO
	return m
}

func (m *Model) MapFromByFieldName(src interface{}) interfaces.IModel {
	if m.err != nil {
		return m
	}
	err := mapper.StructMapByFieldName(src, m.m)
	if err != nil {
		m.err = err
	}
	return m
}

func (m *Model) MapFromByFieldTag(src interface{}) interfaces.IModel {
	//TODO
	return m
}

func (m *Model) ModelError() error {
	return m.err
}
