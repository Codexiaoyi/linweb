package model

import (
	"linweb/interfaces"

	"github.com/Codexiaoyi/go-mapper"
	"github.com/go-playground/validator"
)

type Model struct {
	m interface{}
}

func (m *Model) New(model interface{}) interfaces.IModel {
	return &Model{
		m: model,
	}
}

func (m *Model) Validate() error {
	v := validator.New()
	err := v.Struct(m.m)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) MapToByFieldName(dest interface{}) error {
	err := mapper.StructMapByFieldName(m.m, dest)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) MapToByFieldTag(dest interface{}) error {
	//TODO
	return nil
}

func (m *Model) MapFromByFieldName(src interface{}) error {
	err := mapper.StructMapByFieldName(src, m.m)
	if err != nil {
		return err
	}
	return nil
}

func (m *Model) MapFromByFieldTag(src interface{}) error {
	//TODO
	return nil
}
