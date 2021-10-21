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

package model

import (
	"github.com/Codexiaoyi/linweb/interfaces"

	"github.com/Codexiaoyi/go-mapper"
	"github.com/go-playground/validator"
)

var _ interfaces.IModel = &Model{}

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
