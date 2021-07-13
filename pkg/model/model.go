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

func NewModel(model interface{}) interfaces.IModel {
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

func (m *Model) MapTo(dest interface{}) interfaces.IModel {
	if m.err != nil {
		return m
	}
	err := mapper.StructMapByFieldName(m.m, dest)
	if err != nil {
		m.err = err
	}
	return m
}

func (m *Model) MapBy(src interface{}) interfaces.IModel {
	if m.err != nil {
		return m
	}
	err := mapper.StructMapByFieldName(src, m.m)
	if err != nil {
		m.err = err
	}
	return m
}

func (m *Model) ModelError() string {
	if m.err != nil {
		return m.err.Error()
	}
	return ""
}
